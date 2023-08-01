package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/PavelTabacu/auction/testutil/keeper"
	"github.com/PavelTabacu/auction/testutil/nullify"
	"github.com/PavelTabacu/auction/x/auction/types"
)

func TestAssetQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAsset(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAssetRequest
		response *types.QueryGetAssetResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAssetRequest{Id: msgs[0].Id},
			response: &types.QueryGetAssetResponse{Asset: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetAssetRequest{Id: msgs[1].Id},
			response: &types.QueryGetAssetResponse{Asset: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetAssetRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Asset(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAssetQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAsset(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAssetRequest {
		return &types.QueryAllAssetRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AssetAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Asset), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Asset),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AssetAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Asset), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Asset),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AssetAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Asset),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AssetAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
