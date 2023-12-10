package ante

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// SSITxDecorator ensures that transactions containing SSI messages are combined with messages from other modules.
// This is due to the fixed fees associated with SSI messages, irrespective of message size
type SSITxDecorator struct{}

func NewSSITxDecorator() SSITxDecorator {
	return SSITxDecorator{}
}

func (SSITxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	ssiMsgs, nonSsiMsgs := filterMsgsIntoSSIAndNonSSI(msgs)

	if len(ssiMsgs) != 0 && len(nonSsiMsgs) != 0 {
		return ctx, fmt.Errorf("combining SSI and non-SSI messages in a transaction is not allowed")
	}

	return next(ctx, tx, simulate)
}

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true and transaction with only non-SSI messages
// If fee is high enough or not CheckTx or the transaction consists of only SSI messages, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type MempoolFeeDecorator struct{}

func NewMempoolFeeDecorator() MempoolFeeDecorator {
	return MempoolFeeDecorator{}
}

func (mfd MempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Get list of non SSI messages
	msgs := tx.GetMsgs()
	_, nonSSIMsgs := filterMsgsIntoSSIAndNonSSI(msgs)

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx. SSI messages incur fixed fee regardless of their size
	// and hence any transaction consisting of only SSI messages should skip the following check
	if ctx.IsCheckTx() && !simulate && (len(nonSSIMsgs) > 0) {
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			requiredFees := make(sdk.Coins, len(minGasPrices))

			// Determine the required fees by multiplying each required minimum gas
			// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
			glDec := sdk.NewDec(int64(gas)) /* #nosec G701 */
			for i, gp := range minGasPrices {
				fee := gp.Amount.Mul(glDec)
				requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
			}

			if !feeCoins.IsAnyGTE(requiredFees) {
				return ctx, errors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// ModifiedFeeDecorator is modified version of Cosmos SDK's DeductFeeDecorator implementation which extends it and makes all x/ssi module
// related transactions incur fixed fee cost. Fixed fee is seperate for each x/ssi module transactions and are updated through Governance
// based proposals. Please refer HIP-9 to know more: https://github.com/hypersign-protocol/HIPs/blob/main/HIPs/hip-9.md
type DeductFeeDecorator struct {
	ak             AccountKeeper
	bankKeeper     BankKeeper
	feegrantKeeper FeegrantKeeper
	ssiKeeper      SsiKeeper
}

func NewDeductFeeDecorator(ak AccountKeeper, bk BankKeeper, fk FeegrantKeeper, ifk SsiKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:             ak,
		bankKeeper:     bk,
		feegrantKeeper: fk,
		ssiKeeper:      ifk,
	}
}

func (mfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := mfd.ak.GetModuleAddress(types.FeeCollectorName); addr == nil {
		return ctx, fmt.Errorf("fee collector module account (%s) has not been set", types.FeeCollectorName)
	}

	fee := feeTx.GetFee()
	feePayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter set deduct fee from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if mfd.feegrantKeeper == nil {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest, "fee grants are not enabled")
		} else if !feeGranter.Equals(feePayer) {
			err := mfd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, fee, tx.GetMsgs())
			if err != nil {
				return ctx, errors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}
	deductFeesFromAcc := mfd.ak.GetAccount(ctx, deductFeesFrom)

	// Get the mint module address
	if deductFeesFromAcc == nil {
		return ctx, errors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// Filter SSI messages from Tx messages
	ssiMsgs, _ := filterMsgsIntoSSIAndNonSSI(tx.GetMsgs())

	// If present, calculate the total fee of all fixed-cost x/ssi messages
	if len(ssiMsgs) > 0 {
		fixedSSIFee, err := calculateSSIFeeFromMsgs(ctx, mfd.ssiKeeper, ssiMsgs)
		if err != nil {
			return ctx, err
		}

		// If there is atleast one x/ssi message, check if the fee provided meets the requirement for the fixedSSIFee by asserting if the former
		// is greater than or equal to the latter. If fixedSSIFee is Zero, go ahead with normal fee deduction
		if !fee.IsEqual(fixedSSIFee) {
			errMsg1 := "the transaction consists of x/ssi module based messages which incur fixed cost. "
			errMsg2 := "The fee provided MUST BE equal to the sum of all fixed-fee x/ssi messages. "
			errMsg3 := "To know about the fixed-fee cost of all x/ssi transactions, refer the API endpoint /hypersign-protocol/hidnode/fixedfee . "

			return ctx, errors.Wrapf(
				sdkerrors.ErrInsufficientFee,
				errMsg1+errMsg2+errMsg3+"expected fees to be %v, got %v",
				fixedSSIFee.String(),
				fee.String(),
			)
		}

		// Deduct fixed SSI fee
		err = deductFees(mfd.bankKeeper, ctx, deductFeesFromAcc, fee)
		if err != nil {
			return ctx, err
		}

		events := sdk.Events{
			sdk.NewEvent(
				sdk.EventTypeTx,
				sdk.NewAttribute(sdk.AttributeKeyFee, fixedSSIFee.String()),
				sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
			),
		}
		ctx.EventManager().EmitEvents(events)

		return next(ctx, tx, simulate)
	}

	// Deduct default fees
	if !feeTx.GetFee().IsZero() {
		err := deductFees(mfd.bankKeeper, ctx, deductFeesFromAcc, fee)
		if err != nil {
			return ctx, err
		}
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	}
	ctx.EventManager().EmitEvents(events)

	return next(ctx, tx, simulate)
}

// deductFees deducts fees from the given account.
func deductFees(bankKeeper BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins) error {
	if !fees.IsValid() {
		return errors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}
