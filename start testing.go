import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the redeemInfo
	for _, elem := range genState.RedeemInfoList {
		k.SetRedeemInfo(ctx, elem)
	}
	// Set all the trade
	for _, elem := range genState.TradeList {
		k.SetTrade(ctx, elem)

	t.Run("Cookbook", func(t *testing.T) {
		_, err := util.GenerateAddressWithAccount(ctx, t, net)
		require.NoError(t, err)
		args := []string{"NewUser0", goodPLC}
		_, err = clitestutil.ExecTestCLICmd(ctx, DevCreate(), args)
		assert.Nil(t, err)
		args = []string{"NewUser0", strings.Replace(goodPLC, "0.0.1", "0.0.2", 1)}
		cmd := DevUpdate()
		_, err = clitestutil.ExecTestCLICmd(ctx, cmd, args)
		assert.Nil(t, err)
	}

	// Set all the execution
	for _, elem := range genState.ExecutionList {
		k.SetExecution(ctx, elem)
	}
