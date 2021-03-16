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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// NetworkAPIService implements the server.NetworkAPIServicer interface.
type NetworkAPIService struct {
	network *types.NetworkIdentifier
}

type Peer struct {
	Result []struct {
		Addr string `json:"addr"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    interface{} `json:"id"`
}

// NewNetworkAPIService creates a new instance of a NetworkAPIService.
func NewNetworkAPIService(network *types.NetworkIdentifier) server.NetworkAPIServicer {
	return &NetworkAPIService{
		network: network,
	}
}

// NetworkList implements the /network/list endpoint
func (s *NetworkAPIService) NetworkList(
	ctx context.Context,
	request *types.MetadataRequest,
) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			s.network,
		},
	}, nil
}

// NetworkStatus implements the /network/status endpoint.
func (s *NetworkAPIService) NetworkStatus(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkStatusResponse, *types.Error) {

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

	b_url := fmt.Sprintf("%s/%s/%d", url, "header", height)
	resp, _ := http.Get(b_url)
	resp_body = nil
	if resp != nil {
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			json.NewDecoder(resp.Body).Decode(&resp_body)
		} else {
			fmt.Println("Argh! Broken")
		}
		defer resp.Body.Close()
	}
	curr_block := &types.BlockIdentifier{
		Index: height,
		Hash:  fmt.Sprintf("%v", chain["tip"]),
	}

	// ---------------------------------------------------------------

	// genesis block data --------------------------------------------
	b_url = fmt.Sprintf("%s/%s/%d", url, "header", 0)
	resp, _ = http.Get(b_url)
	var genesis_block map[string]interface{}
	if resp != nil {
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			json.NewDecoder(resp.Body).Decode(&genesis_block)
		} else {
			fmt.Println("Argh! Broken")
		}
		defer resp.Body.Close()
	}

	gen_block := &types.BlockIdentifier{
		Index: 0,
		Hash:  fmt.Sprintf("%v", genesis_block["hash"]),
	}
	// ---------------------------------------------------------------

	// getting peers--------------------------------------------------
	jsonStr := []byte(`{ "method": "getpeerinfo", "params": []  }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var peer Peer
	peers := make([]*types.Peer, 0)
	json.NewDecoder(resp.Body).Decode(&peer)

	for _, v := range peer.Result {
		pr := &types.Peer{PeerID: v.Addr}
		peers = append(peers, pr)
	}
	// -------------------------------------------------------

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: curr_block,
		CurrentBlockTimestamp:  int64(resp_body["time"].(float64)) * 1000,
		GenesisBlockIdentifier: gen_block,
		Peers:                  peers,
	}, nil
}

// NetworkOptions implements the /network/options endpoint.
func (s *NetworkAPIService) NetworkOptions(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkOptionsResponse, *types.Error) {
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.4.0",
			NodeVersion:    "0.0.1",
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     "Success",
					Successful: true,
				},
			},
			OperationTypes: []string{
				"Transfer",
			},
			Errors: []*types.Error{
				{
					Code:      1,
					Message:   "Transaction is required.",
					Retriable: true,
				},
				{
					Code:      2,
					Message:   "Transaction hash is required.",
					Retriable: true,
				},
				{
					Code:      3,
					Message:   "Block is required.",
					Retriable: true,
				},
				{
					Code:      4,
					Message:   "Block height is required.",
					Retriable: true,
				},
				{
					Code:      5,
					Message:   "Block hash mismatch.",
					Retriable: true,
				},
				{
					Code:      6,
					Message:   "Account is required.",
					Retriable: true,
				}, {
					Code:      7,
					Message:   "Address is required.",
					Retriable: true,
				},
				{
					Code:      8,
					Message:   "Network is required.",
					Retriable: true,
				},
				{
					Code:      9,
					Message:   "Options is required.",
					Retriable: true,
				},
				{
					Code:      10,
					Message:   "Invalid network.",
					Retriable: true,
				},
				{
					Code:      11,
					Message:   "Invalid blockchain.",
					Retriable: true,
				},
				{
					Code:      12,
					Message:   "Invalid transaction.",
					Retriable: true,
				},
				{
					Code:      13,
					Message:   "Invalid signed transaction.",
					Retriable: true,
				},
				{
					Code:      14,
					Message:   "Block not found.",
					Retriable: true,
				},
				{
					Code:      15,
					Message:   "Transaction not found.",
					Retriable: true,
				},
				{
					Code:      16,
					Message:   "Coin view not found.",
					Retriable: true,
				},
				{
					Code:      17,
					Message:   "Error relaying transaction.",
					Retriable: false,
				},
				{
					Code:      18,
					Message:   "Query not supported",
					Retriable: false,
				},
				{
					Code:      32,
					Message:   "Unknown error.",
					Retriable: false,
				},
			},
		},
	}, nil
}
