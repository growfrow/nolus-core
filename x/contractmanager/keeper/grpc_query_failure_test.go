package keeper_test

import (
	"strconv"
	"testing"

	"github.com/Nolus-Protocol/nolus-core/testutil/common/nullify"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/Nolus-Protocol/nolus-core/testutil/contractmanager/keeper"
	"github.com/Nolus-Protocol/nolus-core/x/contractmanager/keeper"
	"github.com/Nolus-Protocol/nolus-core/x/contractmanager/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFailureQuerySingle(t *testing.T) {
	k, ctx := keepertest.ContractManagerKeeper(t, nil)
	msgs := createNFailure(k, ctx, 2, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryFailuresByAddressRequest
		response *types.QueryAddressFailuresResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryFailuresByAddressRequest{
				Address: msgs[0][0].Address,
			},
			response: &types.QueryAddressFailuresResponse{Failures: msgs[0], Pagination: &query.PageResponse{Total: 2}},
		},
		{
			desc: "Second",
			request: &types.QueryFailuresByAddressRequest{
				Address: msgs[1][0].Address,
			},
			response: &types.QueryAddressFailuresResponse{Failures: msgs[1], Pagination: &query.PageResponse{Total: 2}},
		},
		{
			desc: "KeyIsAbsent",
			request: &types.QueryFailuresByAddressRequest{
				Address: "nolus1f6cu6ypvpyh0p8d7pqnps2pduj87hda5t9v4mqrc8ra67xp28uwq4f4ysz",
			},
			response: &types.QueryAddressFailuresResponse{Failures: []types.Failure{}, Pagination: &query.PageResponse{Total: 0}},
		},
		{
			desc: "InvalidAddress",
			request: &types.QueryFailuresByAddressRequest{
				Address: "wrong_address",
			},
			err: status.Error(codes.InvalidArgument, "failed to parse address: wrong_address"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.AddressFailures(ctx, tc.request)
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

func TestFailureQueryPaginated(t *testing.T) {
	k, ctx := keepertest.ContractManagerKeeper(t, nil)
	msgs := createNFailure(k, ctx, 5, 3)
	flattenItems := flattenFailures(msgs)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryFailuresRequest {
		return &types.QueryFailuresRequest{
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
		for i := 0; i < len(flattenItems); i += step {
			resp, err := k.Failures(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Failures), step)
			require.Subset(t,
				nullify.Fill(flattenItems),
				nullify.Fill(resp.Failures),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(flattenItems); i += step {
			resp, err := k.Failures(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Failures), step)
			require.Subset(t,
				nullify.Fill(flattenItems),
				nullify.Fill(resp.Failures),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.Failures(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(flattenItems), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(flattenItems),
			nullify.Fill(resp.Failures),
		)
	})
	t.Run("MoreThanLimit", func(t *testing.T) {
		_, err := k.Failures(ctx, request(nil, 0, keeper.FailuresQueryMaxLimit+1, true))
		require.ErrorContains(t, err, "limit is more than maximum allowed")
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.Failures(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
