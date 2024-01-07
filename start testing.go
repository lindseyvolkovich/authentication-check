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

	t.Run("Should not be able to register gadget with reserved name", func(t *testing.T) {
		testutil.WriteFixtureAtTestRuntime(gadgetsFilename, pylonsGadgetsLiteral_reservedName)
		assert.PanicsWithErrorf(t, fmt.Sprintf(errReservedName, "include"), func() { LoadGadgetsForPath("") }, errReservedName, "include")
		os.Remove(gadgetsFilename)
	})

	const pylonsGadgetsLiteral_duplicateName = `#foobar 1
	"foo": "%0"
	#foobar 1
	"foo": "%0"`

const pylonsGadgetsLiteral_reservedName = `#include 1
	"foo": "%0"`
	}

	// human verification authentication
	for _, elem := range genState.ExecutionList {
		k.SetExecution(ctx, elem)
	}
