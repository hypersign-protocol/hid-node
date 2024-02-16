package tests

// import (
// 	"testing"

// 	testutil "github.com/hypersign-protocol/hid-node/testutil/keeper"
// 	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
// 	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
// 	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"

// 	// "github.com/hypersign-protocol/hid-node/x/ssi/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// func TestCreateDIDUsingEd25519KeyPair(t *testing.T) {
// 	t.Log("Running test for Valid Create DID Tx")
// 	k, ctx := testutil.SsiKeeper(t)
// 	msgServer := keeper.NewMsgServerImpl(*k)
// 	goCtx := sdk.WrapSDKContext(ctx)

// 	k.SetChainNamespace(&ctx, chainNamespace)

// 	keyPair1 := testcrypto.GenerateEd25519KeyPair()
// 	rpcElements := GenerateDidDocumentRPCElements(keyPair1, []*testssi.SsiDocSigningElements{})
// 	t.Logf("Registering DID with DID Id: %s", rpcElements.DidDocument.GetId())

// 	msgRegisterDID := rpcElements

// 	_, err := msgServer.RegisterDID(goCtx, msgRegisterDID)
// 	if err != nil {
// 		t.Error("DID Registeration Failed")
// 		t.Error(err)
// 		t.FailNow()
// 	}
// 	t.Log("Did Registeration Successful")

// 	t.Log("Create DID Tx Test Completed")
// }
