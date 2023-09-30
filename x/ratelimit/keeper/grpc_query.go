package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	"github.com/notional-labs/centauri/v5/x/ratelimit/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

// NewQueryServer returns an implementation of the QueryServer
// for the provided Keeper.
func NewQueryServer(k Keeper) types.QueryServer {
	return queryServer{Keeper: k}
}

// AllRateLimits queries all rate limits.
func (q queryServer) AllRateLimits(c context.Context, req *types.QueryAllRateLimitsRequest) (*types.QueryAllRateLimitsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	rateLimits := q.GetAllRateLimits(ctx)
	return &types.QueryAllRateLimitsResponse{RateLimits: rateLimits}, nil
}

// RateLimit queries a rate limit by denom and channel id.
func (q queryServer) RateLimit(c context.Context, req *types.QueryRateLimitRequest) (*types.QueryRateLimitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	rateLimit, found := q.GetRateLimit(ctx, req.Denom, req.ChannelID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "rate limit by denom %s and channel id %s not found", req.Denom, req.ChannelID)
	}
	return &types.QueryRateLimitResponse{RateLimit: &rateLimit}, nil
}

// RateLimitsByChainID queries all rate limits for a given chain.
func (q queryServer) RateLimitsByChainID(c context.Context, req *types.QueryRateLimitsByChainIDRequest) (*types.QueryRateLimitsByChainIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	chainId := req.ChainId

	rateLimits := []types.RateLimit{}
	for _, rateLimit := range q.GetAllRateLimits(ctx) {
		_, clientState, err := q.channelKeeper.GetChannelClientState(ctx, transfertypes.PortID, rateLimit.Path.ChannelID)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "unable to fetch client state by port id %s and channel id: %s", transfertypes.PortID, rateLimit.Path.ChannelID)
		}

		client, ok := clientState.(*ibctmtypes.ClientState)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "invalid client state")
		}

		// Append the rate limit when it matches with the requested chain id
		if client.ChainId == chainId {
			rateLimits = append(rateLimits, rateLimit)
		}
	}

	return &types.QueryRateLimitsByChainIDResponse{RateLimits: rateLimits}, nil
}

// RateLimitsByChannelID queries all rate limits for a given channel.
func (q queryServer) RateLimitsByChannelID(c context.Context, req *types.QueryRateLimitsByChannelIDRequest) (*types.QueryRateLimitsByChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	rateLimits := []types.RateLimit{}
	for _, rateLimit := range q.GetAllRateLimits(ctx) {
		// If the channel ID matches, add the rate limit to the returned list
		if rateLimit.Path.ChannelID == req.ChannelID {
			rateLimits = append(rateLimits, rateLimit)
		}
	}

	return &types.QueryRateLimitsByChannelIDResponse{RateLimits: rateLimits}, nil
}

// AllWhitelistedAddresses queries all whitelisted addresses.
func (q queryServer) AllWhitelistedAddresses(c context.Context, req *types.QueryAllWhitelistedAddressesRequest) (*types.QueryAllWhitelistedAddressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	whitelistedAddresses := q.GetAllWhitelistedAddressPairs(ctx)
	return &types.QueryAllWhitelistedAddressesResponse{AddressPairs: whitelistedAddresses}, nil
}
