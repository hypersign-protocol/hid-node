package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	identityfeetypes "github.com/hypersign-protocol/hid-node/x/identityfee/types"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
)

type DeductFeeDecorator struct {
	ak                AccountKeeper
	bankKeeper        BankKeeper
	feegrantKeeper    FeegrantKeeper
	identityFeeKeeper IdentityFeeKeeper
}

// ModifiedFeeDecorator extends NewDeductFeeDecorator by making all x/ssi module related transactions incur fixed fee cost.
// Fixed fee is seperate for each x/ssi module transactions and are updated through Governance based proposals. Please refer
// HIP-9 to know more: https://github.com/hypersign-protocol/HIPs/blob/main/HIPs/hip-9.md
func NewDeductFeeDecorator(ak AccountKeeper, bk BankKeeper, fk FeegrantKeeper, ifk IdentityFeeKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:                ak,
		bankKeeper:        bk,
		feegrantKeeper:    fk,
		identityFeeKeeper: ifk,
	}
}

func (mfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
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
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee grants are not enabled")
		} else if !feeGranter.Equals(feePayer) {
			err := mfd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, fee, tx.GetMsgs())
			if err != nil {
				return ctx, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}
	deductFeesFromAcc := mfd.ak.GetAccount(ctx, deductFeesFrom)

	// Get the mint module address
	if deductFeesFromAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", deductFeesFrom)
	}

	// If present, calculate the total fee of all fixed-cost x/ssi messages
	if isSSIMsgPresentInTx(feeTx) {
		var fixedSSIFee sdk.Coins = sdk.NewCoins(sdk.NewCoin("uhid", sdk.NewInt(0)))

		for _, msg := range tx.GetMsgs() {
			switch msg.(type) {
			case *ssitypes.MsgCreateDID:
				createDidFee := mfd.identityFeeKeeper.GetFeeParams(ctx, identityfeetypes.ParamStoreKeyCreateDidFee)
				fixedSSIFee = fixedSSIFee.Add(createDidFee)
			case *ssitypes.MsgUpdateDID:
				updateDidFee := mfd.identityFeeKeeper.GetFeeParams(ctx, identityfeetypes.ParamStoreKeyUpdateDidFee)
				fixedSSIFee = fixedSSIFee.Add(updateDidFee)
			case *ssitypes.MsgDeactivateDID:
				deactivateDidFee := mfd.identityFeeKeeper.GetFeeParams(ctx, identityfeetypes.ParamStoreKeyDeactivateDidFee)
				fixedSSIFee = fixedSSIFee.Add(deactivateDidFee)
			case *ssitypes.MsgCreateSchema:
				createSchemaFee := mfd.identityFeeKeeper.GetFeeParams(ctx, identityfeetypes.ParamStoreKeyCreateSchemaFee)
				fixedSSIFee = fixedSSIFee.Add(createSchemaFee)
			case *ssitypes.MsgRegisterCredentialStatus:
				registerCredentialStatusFee := mfd.identityFeeKeeper.GetFeeParams(ctx, identityfeetypes.ParamStoreKeyRegisterCredentialStatusFee)
				fixedSSIFee = fixedSSIFee.Add(registerCredentialStatusFee)
			}
		}

		// If there is atleast one x/ssi message, check if the fee provided meets the requirement for the fixedSSIFee by asserting if the former
		// is greater than or equal to the latter. If fixedSSIFee is Zero, go ahead with normal fee deduction
		if !fee.IsEqual(fixedSSIFee) {
			errMsg1 := "the transaction consists of x/ssi module based messages which incurs fixed cost. "
			errMsg2 := "The fee provided MUST BE equal to the required fees which is the sum of all fixed-fee x/ssi. "
			errMsg3 := "To know about the fixed-fee cost of all x/ssi transactions, refer to the API endpoint /hypersign-protocol/hidnode/identityfee . "

			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrInsufficientFee,
				errMsg1+errMsg2+errMsg3+"expected fees to be %v, got %v",
				fixedSSIFee.String(),
				fee.String(),
			)
		}

		// Deduct fixed SSI fee
		err := deductFees(mfd.bankKeeper, ctx, deductFeesFromAcc, fee)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}

// isSSIMsgPresentInTx checks if there is any SSI message present in the transaction
func isSSIMsgPresentInTx(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if isSSIMsg(msg) {
			return true
		}
	}

	return false
}

// isNonSSIMsgPresentInTx checks if there is any non-SSI message present in the transaction
func isNonSSIMsgPresentInTx(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if !isSSIMsg(msg) {
			return true
		}
	}

	return false
}

// isSSIMsg checks if the message is of SSI type or not
func isSSIMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case *ssitypes.MsgCreateDID:
		return true
	case *ssitypes.MsgUpdateDID:
		return true
	case *ssitypes.MsgDeactivateDID:
		return true
	case *ssitypes.MsgCreateSchema:
		return true
	case *ssitypes.MsgRegisterCredentialStatus:
		return true
	default:
		return false
	}
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
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx. SSI messages incur fixed fee regardless of their size
	// and hence any transaction consisting of only SSI messages should skip the following check
	if ctx.IsCheckTx() && !simulate && !isSSIMsgPresentInTx(feeTx) {
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			requiredFees := make(sdk.Coins, len(minGasPrices))

			// Determine the required fees by multiplying each required minimum gas
			// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
			glDec := sdk.NewDec(int64(gas))
			for i, gp := range minGasPrices {
				fee := gp.Amount.Mul(glDec)
				requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
			}

			if !feeCoins.IsAnyGTE(requiredFees) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// SSITxDecorator ensures that any transaction which contains SSI messages should not have messages of other modules, since fees for
// SSI messages are fixed regardless of the size of message, they should not be mixed other module's messages.
type SSITxDecorator struct{}

func NewSSITxDecorator() SSITxDecorator {
	return SSITxDecorator{}
}

func (ssifd SSITxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if isNonSSIMsgPresentInTx(tx) && isSSIMsgPresentInTx(tx) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "ssi transaction messages cannot be grouped with other module messages")
	}

	return next(ctx, tx, simulate)
}
