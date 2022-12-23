package keeper_test

import (
	"testing"

	testkeeper "chimba/testutil/keeper"
	"chimba/x/chimba/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChimbaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
