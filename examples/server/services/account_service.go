// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// Account implements the /Account endpoint.
func (s *AccountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	// -----------------------------------------------------------
	// check port

	var resp_body map[string]interface{}
	var url string
	port_mapping := map[string]string{"13037": "testnet", "14037": "testnet", "12037": "main"}
	for k := range port_mapping {
		hsd_node := "http://127.0.0.1"
		url = fmt.Sprintf("%s:%s", hsd_node, k)
		resp, _ := http.Get(url)
		if resp != nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
				json.NewDecoder(resp.Body).Decode(&resp_body)
				defer resp.Body.Close()
				break
			} else {
				fmt.Println("Argh! Broken")
			}
			defer resp.Body.Close()
		}
	}
	// ------------------------------------------------------------

	// current block data -----------------------------------------
	chain := resp_body["chain"].(map[string]interface{})
	height, _ := strconv.ParseInt(fmt.Sprintf("%v", chain["height"]), 10, 64)
	curr_block := &types.BlockIdentifier{
		Index: height,
		Hash:  fmt.Sprintf("%v", chain["tip"]),
	}
	return &types.AccountBalanceResponse{
		BlockIdentifier: curr_block,
		Balances:        nil,
	}, nil
}

func (s *AccountAPIService) AccountCoins(
	ctx context.Context,
	request *types.AccountCoinsRequest,
) (*types.AccountCoinsResponse, *types.Error) {

	return nil, nil
}
