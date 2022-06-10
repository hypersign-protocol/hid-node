package tests

import (
	"testing"

	testkeeper "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateSchema(t *testing.T) {
	t.Log("Running test for Valid Create Schmea Tx")
	k, ctx := testkeeper.SsiKeeper(t)
	msgServ := keeper.NewMsgServerImpl(*k)
	goCtx :=  sdk.WrapSDKContext(ctx)
	
	t.Logf("Registering DID with Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures: DidDocumentValidSignInfo,
		Creator: Creator,
	}

	_, err := msgServ.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("DID Registered Successfully")

	t.Log("Registering Schema")
	msgCreateSchema := &types.MsgCreateSchema{
		Schema: ValidSchemaDocument,
		Signatures: SchemaValidSignInfo,
		Creator: Creator,
	}

	_, errCreateSchema := msgServ.CreateSchema(goCtx, msgCreateSchema)
	if errCreateSchema != nil {
		t.Error("Schema Registeration Failed")
		t.Error(errCreateSchema)
		t.FailNow()
	} 
	t.Log("Schema Registered Successfully")

	t.Log("Create Schema Tx Test Completed")
}