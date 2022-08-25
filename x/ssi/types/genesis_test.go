package types_test

import (
	"testing"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid did namespace",
			genState: &types.GenesisState{
				ChainNamespace: "devnet",
			},
			valid: true,
		},
		{
			desc:     "invalid did namespace of length more than 10",
			genState: &types.GenesisState{
				ChainNamespace: "abracadabra123",
			},
			valid: false,
		},
		{
			desc:     "invalid did namespace containing whitespaces",
			genState: &types.GenesisState{
				ChainNamespace: "abracadabra	123",
			},
			valid: false,
		},
		{
			desc:     "invalid did namespace containing underscore",
			genState: &types.GenesisState{
				ChainNamespace: "xyz_123",
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
