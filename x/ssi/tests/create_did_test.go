package tests

import (
	"testing"

	testkeeper "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateDID(t *testing.T) {
	t.Log("Running test for Valid Create DID Tx")
	k, ctx := testkeeper.SsiKeeper(t)
	msgServ := keeper.NewMsgServerImpl(*k)
	
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocumet,
		Signatures: ValidSignInfo,
		Creator: Creator,
	}

	goCtx :=  sdk.WrapSDKContext(ctx)
	_, err := msgServ.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error(err)
	}

	t.Log("Create DID Tx Test Completed")
}