package tests

import (
	"testing"

	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateDID(t *testing.T) {
	t.Log("Running test for Valid Create DID Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx :=  sdk.WrapSDKContext(ctx)
	
	t.Logf("Registering DID with DID Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures: DidDocumentValidSignInfo,
		Creator: Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Did Registeration Successful")

	t.Log("Create DID Tx Test Completed")
}