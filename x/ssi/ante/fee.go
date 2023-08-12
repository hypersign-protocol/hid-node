package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// getFeeForSSIMsg returns fee for the input SSI message
func getFeeForSSIMsg(ctx sdk.Context, msg sdk.Msg, ssiKeeper SsiKeeper) sdk.Coin {
	switch msg.(type) {
	case *ssitypes.MsgCreateDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyCreateDidFee)
	case *ssitypes.MsgUpdateDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyUpdateDidFee)
	case *ssitypes.MsgDeactivateDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyDeactivateDidFee)
	case *ssitypes.MsgCreateSchema:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyCreateSchemaFee)
	case *ssitypes.MsgRegisterCredentialStatus:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyRegisterCredentialStatusFee)
	default:
		return sdk.NewCoin("uhid", sdk.NewInt(0))
	}
}

// calculateSSIFeeFromMsgs calculates the total SSI fixed fee from messages
func calculateSSIFeeFromMsgs(ctx sdk.Context, ssiKeeper SsiKeeper, msgs []SSIMsg) (sdk.Coins, error) {
	var totalFee sdk.Coins = sdk.NewCoins(sdk.NewCoin("uhid", sdk.NewInt(0)))

	for _, msg := range msgs {
		msgFee := getFeeForSSIMsg(ctx, msg, ssiKeeper)
		totalFee = totalFee.Add(msgFee)
	}

	return totalFee, nil
}
