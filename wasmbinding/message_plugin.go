package wasmbinding

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/hypersign-protocol/hid-node/wasmbinding/bindings"
	ssiKeeper "github.com/hypersign-protocol/hid-node/x/ssi/keeper"
)

func CustomMessageDecorator(
	ssiKeeper *ssiKeeper.Keeper,
) func(messenger wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			SsiKeeper: *ssiKeeper,
			Wrapped:   old,
		}
	}
}

type CustomMessenger struct {
	SsiKeeper ssiKeeper.Keeper
	Wrapped   wasmkeeper.Messenger
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var contractMsg bindings.SsiContractMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "failed to decode incoming custom cosmos message")
		}

		if contractMsg.SetBlockchainAccountId != nil {
			return m.SetBlockchainAccountIdFunc(ctx, contractMsg.SetBlockchainAccountId.BlockchainAccountId)
		}
	}

	return m.Wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// This is more like a dummy implementation, as currently there is no way to query blockchainAccountId store from outside
func (m *CustomMessenger) SetBlockchainAccountIdFunc(ctx sdk.Context, bcId string) ([]sdk.Event, [][]byte, error) {
	if bcId == "" {
		return nil, nil, fmt.Errorf("blockchain account id cannot be empty")
	}

	m.SsiKeeper.SetBlockchainAddressInStore(&ctx, bcId, "did:hid:sampledid")

	var dummyReturnMsg []byte = []byte("wasm ssi execute interaction done")

	return nil, [][]byte{dummyReturnMsg}, nil
}
