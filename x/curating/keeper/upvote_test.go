package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/testdata"
	"github.com/stretchr/testify/require"
)

func TestCreateUpvote(t *testing.T) {
	_, app, ctx := testdata.CreateTestInput()

	postID := "500"
	vendorID := uint32(1)
	deposit := sdk.NewInt64Coin("ufuel", 1000000)
	addrs := testdata.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))

	err := app.CuratingKeeper.CreateUpvote(ctx, vendorID, postID, addrs[0], addrs[0], 5, deposit)
	require.NoError(t, err)

	_, found := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "post should be found")

	upvote, found := app.CuratingKeeper.GetUpvote(ctx, vendorID, postID, addrs[0])
	require.True(t, found, "upvote should be found")

	require.Equal(t, "25000000ufuel", upvote.VoteAmount.String())
}
