package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetParams(ctx sdk.Context) (params types.Params)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	SetAccount(ctx sdk.Context, acc types.AccountI)
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type FeegrantKeeper interface {
	UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error
}

type IdentityFeeKeeper interface {
	GetFeeParams(ctx sdk.Context, identityfeeParamStoreKey []byte) sdk.Coin
}
