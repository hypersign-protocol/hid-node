package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// getFeeForSSIMsg returns fee for the input SSI message
func getFeeForSSIMsg(ctx sdk.Context, msg sdk.Msg, ssiKeeper SsiKeeper) sdk.Coin {
	switch msg.(type) {
	case *ssitypes.MsgRegisterDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyRegisterDidFee)
	case *ssitypes.MsgUpdateDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyUpdateDidFee)
	case *ssitypes.MsgDeactivateDID:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyDeactivateDidFee)
	case *ssitypes.MsgRegisterCredentialSchema:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyRegisterCredentialSchemaFee)
	case *ssitypes.MsgUpdateCredentialSchema:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyUpdateCredentialSchemaFee)
	case *ssitypes.MsgRegisterCredentialStatus:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyRegisterCredentialStatusFee)
	case *ssitypes.MsgUpdateCredentialStatus:
		return ssiKeeper.GetFeeParams(ctx, ssitypes.ParamStoreKeyUpdateCredentialStatusFee)
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
