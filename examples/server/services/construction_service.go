package services

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// ConstructionAPIService implements the server.ConstructionAPIServicer interface.
type ConstructionAPIService struct {
	network *types.NetworkIdentifier
}

// NewConstructionAPIService creates a new instance of a ConstructionAPIService.
func NewConstructionAPIService(network *types.NetworkIdentifier) server.ConstructionAPIServicer {
	return &ConstructionAPIService{
		network: network,
	}
}

// ConstructionCombine implements the /construction/combine endpoint.
func (s *ConstructionAPIService) ConstructionCombine(
	ctx context.Context,
	request *types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	return nil, nil
}

// ConstructionDerive implements the /construction/derive endpoint.
func (s *ConstructionAPIService) ConstructionDerive(
	ctx context.Context,
	request *types.ConstructionDeriveRequest,
) (*types.ConstructionDeriveResponse, *types.Error) {
	return nil, nil
}

// ConstructionHash implements the /construction/hash endpoint.
func (s *ConstructionAPIService) ConstructionHash(
	ctx context.Context,
	request *types.ConstructionHashRequest,
) (*types.TransactionIdentifierResponse, *types.Error) {
	return nil, nil
}

// ConstructionMetadata implements the /construction/metadata endpoint.
func (s *ConstructionAPIService) ConstructionMetadata(
	ctx context.Context,
	request *types.ConstructionMetadataRequest,
) (*types.ConstructionMetadataResponse, *types.Error) {
	return nil, nil
}

// ConstructionParse implements the /ConstructionParse endpoint.
func (s *ConstructionAPIService) ConstructionParse(
	ctx context.Context,
	request *types.ConstructionParseRequest,
) (*types.ConstructionParseResponse, *types.Error) {
	status := new(string)
	*status = "Reverted"
	return &types.ConstructionParseResponse{
		Operations: []*types.Operation{
			{
				OperationIdentifier: &types.OperationIdentifier{
					Index:        1,
					NetworkIndex: nil,
				},
				RelatedOperations: []*types.OperationIdentifier{
					{
						Index:        0,
						NetworkIndex: nil,
					},
				},
				Type:   "PAYMENT",
				Status: status,
				Account: &types.AccountIdentifier{
					Address:    "sender address 666",
					SubAccount: nil,
					Metadata:   nil,
				},
				Amount: &types.Amount{
					Value: "1000",
					Currency: &types.Currency{
						Symbol:   "ROS",
						Decimals: 2,
						Metadata: nil,
					},
					Metadata: nil,
				},
				CoinChange: &types.CoinChange{
					CoinIdentifier: &types.CoinIdentifier{
						Identifier: "Coin 1",
					},
					CoinAction: "coin_created",
				},
				Metadata: nil,
			},
		},
		AccountIdentifierSigners: []*types.AccountIdentifier{
			{
				Address:    "sender address 666",
				SubAccount: nil,
				Metadata:   nil,
			},
		},
		Metadata: nil,
	}, nil
}

// ConstructionPayloads implements the /construction/payloads endpoint.
func (s *ConstructionAPIService) ConstructionPayloads(
	ctx context.Context,
	request *types.ConstructionPayloadsRequest,
) (*types.ConstructionPayloadsResponse, *types.Error) {
	return &types.ConstructionPayloadsResponse{
		UnsignedTransaction: "transaction 100",
		Payloads: []*types.SigningPayload{
			{
				AccountIdentifier: &types.AccountIdentifier{
					Address:    "sender address 666",
					SubAccount: nil,
					Metadata:   nil,
				},
				Bytes:         []byte{0xe0}, //byte randomly set
				SignatureType: "ecdsa",
			},
		},
	}, nil
}

// ConstructionPreprocess implements the /construction/preprocess endpoint.
func (s *ConstructionAPIService) ConstructionPreprocess(
	ctx context.Context,
	request *types.ConstructionPreprocessRequest,
) (*types.ConstructionPreprocessResponse, *types.Error) {
	return nil, nil
}

// ConstructionSubmit implements the /construction/submit endpoint.
func (s *ConstructionAPIService) ConstructionSubmit(
	ctx context.Context,
	request *types.ConstructionSubmitRequest,
) (*types.TransactionIdentifierResponse, *types.Error) {
	return nil, nil
}
