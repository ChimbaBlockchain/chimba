package chimba_test

import (
	"testing"

	keepertest "chimba/testutil/keeper"
	"chimba/testutil/nullify"
	"chimba/x/chimba"
	"chimba/x/chimba/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChimbaKeeper(t)
	chimba.InitGenesis(ctx, *k, genesisState)
	got := chimba.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
