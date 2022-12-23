package keeper_test

import (
	"context"
	"testing"

	keepertest "chimba/testutil/keeper"
	"chimba/x/chimba/keeper"
	"chimba/x/chimba/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChimbaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
