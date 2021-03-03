package services

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// AccountAPIService implements the server.AccountAPIServicer interface.
type AccountAPIService struct {
	network *types.NetworkIdentifier
}

// NewAccountAPIService creates a new instance of a AccountAPIService.
func NewAccountAPIService(network *types.NetworkIdentifier) server.AccountAPIServicer {
	return &AccountAPIService{
		network: network,
	}
}

// AccountBalance implements the account/balance endpoint
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	return &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "Block 1000",
		},
		Balances: []*types.Amount{
			{
				Value: "1000",
				Currency: &types.Currency{
					Symbol:   "ROS",
					Decimals: 6,
				},
			},
		},
		Metadata: map[string]interface{}{},
	}, nil
}

// AccountCoins implements the account/coins endpoint
func (s *AccountAPIService) AccountCoins(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {
	return &types.AccountCoinsResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: 1000,
			Hash:  "Block 1000",
		},
		Coins: []*types.Coin{
			{
				CoinIdentifier: &types.CoinIdentifier{
					Identifier: "Coin 6798",
				},
				Amount: &types.Amount{
					Value: "25",
					Currency: &types.Currency{
						Symbol:   "ROS",
						Decimals: 6,
						Metadata: map[string]interface{}{},
					},
					Metadata: map[string]interface{}{},
				},
			},
		},
		Metadata: map[string]interface{}{},
	}, nil
}
