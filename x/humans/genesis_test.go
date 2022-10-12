package humans_test

import (
	"testing"

	keepertest "github.com/humansdotai/humans/testutil/keeper"
	"github.com/humansdotai/humans/testutil/nullify"
	"github.com/humansdotai/humans/x/humans"
	"github.com/humansdotai/humans/x/humans/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FeeBalanceList: []types.FeeBalance{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HumansKeeper(t)
	humans.InitGenesis(ctx, *k, genesisState)
	got := humans.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FeeBalanceList, got.FeeBalanceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
