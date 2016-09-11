// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package web3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/alanchchen/web3go/common"
	"github.com/alanchchen/web3go/rpc"
)

// Eth ...
type Eth interface {
	ProtocolVersion() (string, error)
	Syncing() (common.SyncStatus, error)
	Coinbase() (common.Address, error)
	Mining() (bool, error)
	HashRate() (uint64, error)
	GasPrice() (*big.Int, error)
	Accounts() ([]common.Address, error)
	BlockNumber() (*big.Int, error)
	GetBalance(address common.Address, quantity string) (*big.Int, error)
	GetStorageAt(address common.Address, position uint64, quantity string) (uint64, error)
	GetTransactionCount(address common.Address, quantity string) (*big.Int, error)
	GetBlockTransactionCountByHash(hash common.Hash) (*big.Int, error)
	GetBlockTransactionCountByNumber(quantity string) (*big.Int, error)
	GetUncleCountByBlockHash(hash common.Hash) (*big.Int, error)
	GetUncleCountByBlockNumber(quantity string) (*big.Int, error)
	GetCode(address common.Address, quantity string) ([]byte, error)
	Sign(address common.Address, data []byte) ([]byte, error)
	SendTransaction(tx *common.TransactionRequest) (common.Hash, error)
	SendRawTransaction(tx []byte) (common.Hash, error)
	Call(tx *common.TransactionRequest, quantity string) ([]byte, error)
	EstimateGas(tx *common.TransactionRequest, quantity string) (*big.Int, error)
	GetBlockByHash(hash common.Hash, full bool) (*common.Block, error)
	GetBlockByNumber(quantity string, full bool) (*common.Block, error)
	GetTransactionByHash(hash common.Hash) (*common.Transaction, error)
	GetTransactionByBlockHashAndIndex(hash common.Hash, index uint64) (*common.Transaction, error)
	GetTransactionByBlockNumberAndIndex(quantity string, index uint64) (*common.Transaction, error)
	GetTransactionReceipt(hash common.Hash) (*common.TransactionReceipt, error)
	GetUncleByBlockHashAndIndex(hash common.Hash, index uint64) (*common.Block, error)
	GetUncleByBlockNumberAndIndex(quantity string, index uint64) (*common.Block, error)
	GetCompilers() ([]string, error)
	// GompileLLL
	// CompileSolidity
	// CompileSerpent
	NewFilter(option *FilterOption) (Filter, error)
	NewBlockFilter() (Filter, error)
	NewPendingTransactionFilter() (Filter, error)
	UninstallFilter(filter Filter) (bool, error)
	GetFilterChanges(filter Filter) ([]interface{}, error)
	GetFilterLogs(filter Filter) ([]interface{}, error)
	GetLogs(filter Filter) ([]interface{}, error)
	GetWork() (common.Hash, common.Hash, common.Hash, error)
	SubmitWork(nonce uint64, header common.Hash, mixDigest common.Hash) (bool, error)
	// SubmitHashrate
}

// EthAPI ...
type EthAPI struct {
	rpc            rpc.RPC
	requestManager *requestManager
}

// NewEthAPI ...
func newEthAPI(requestManager *requestManager) Eth {
	return &EthAPI{requestManager: requestManager}
}

// ProtocolVersion returns the current ethereum protocol version.
func (eth *EthAPI) ProtocolVersion() (string, error) {
	req := eth.requestManager.newRequest("eth_protocolVersion")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return "", err
	}

	if resp.Error() != nil {
		return "", resp.Error()
	}

	return resp.Get("result").(string), nil
}

// Syncing returns true with an object with data about the sync status or false
// with nil.
func (eth *EthAPI) Syncing() (common.SyncStatus, error) {
	req := eth.requestManager.newRequest("eth_syncing")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return common.SyncStatus{
			Result: false,
		}, err
	}

	if resp.Error() != nil {
		return common.SyncStatus{
			Result: false,
		}, resp.Error()
	}

	result := resp.Get("result")
	switch result.(type) {
	case bool:
		return common.SyncStatus{
			Result: false,
		}, nil
	default:
		var err error
		var resultBlob []byte
		r := common.SyncStatus{
			Result: true,
		}
		if resultBlob, err = json.Marshal(result); err == nil {
			if err = json.Unmarshal(resultBlob, &r); err == nil {
				return r, nil
			}
		}

		return common.SyncStatus{
			Result: false,
		}, err
	}
}

// Coinbase returns the client coinbase address.
func (eth *EthAPI) Coinbase() (addr common.Address, err error) {
	req := eth.requestManager.newRequest("eth_coinbase")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return common.NewAddress(nil), err
	}

	if resp.Error() != nil {
		return common.NewAddress(nil), resp.Error()
	}

	return common.StringToAddress(resp.Get("result").(string)), nil
}

// Mining returns true if client is actively mining new blocks.
func (eth *EthAPI) Mining() (bool, error) {
	req := eth.requestManager.newRequest("eth_mining")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return false, err
	}

	if resp.Error() != nil {
		return false, resp.Error()
	}

	return resp.Get("result").(bool), nil
}

// HashRate returns the number of hashes per second that the node is mining
// with.
func (eth *EthAPI) HashRate() (uint64, error) {
	req := eth.requestManager.newRequest("eth_hashrate")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return 0, err
	}

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	result, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GasPrice returns the current price per gas in wei.
func (eth *EthAPI) GasPrice() (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_gasPrice")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// Accounts returns a list of addresses owned by client.
func (eth *EthAPI) Accounts() (addrs []common.Address, err error) {
	req := eth.requestManager.newRequest("eth_accounts")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	results := resp.Get("result").([]interface{})
	for _, r := range results {
		addrs = append(addrs, common.StringToAddress(r.(string)))
	}
	return addrs, nil
}

// BlockNumber returns the number of most recent block.
func (eth *EthAPI) BlockNumber() (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_blockNumber")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetBalance returns the balance of the account of given address.
func (eth *EthAPI) GetBalance(address common.Address, quantity string) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getBalance")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetStorageAt returns the value from a storage position at a given address.
func (eth *EthAPI) GetStorageAt(address common.Address, position uint64, quantity string) (uint64, error) {
	req := eth.requestManager.newRequest("eth_getStorageAt")
	req.Set("params", []string{address.String(), fmt.Sprintf("%v", position), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return 0, err
	}

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	result, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetTransactionCount returns the number of transactions sent from an address.
func (eth *EthAPI) GetTransactionCount(address common.Address, quantity string) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getTransactionCount")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetBlockTransactionCountByHash returns the number of transactions in a block
// from a block matching the given block hash.
func (eth *EthAPI) GetBlockTransactionCountByHash(hash common.Hash) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in a
// block from a block matching the given block number.
func (eth *EthAPI) GetBlockTransactionCountByNumber(quantity string) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetUncleCountByBlockHash returns the number of uncles in a block from a block
// matching the given block hash.
func (eth *EthAPI) GetUncleCountByBlockHash(hash common.Hash) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetUncleCountByBlockNumber returns the number of uncles in a block from a
// block matching the given block number.
func (eth *EthAPI) GetUncleCountByBlockNumber(quantity string) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetCode returns code at a given address.
func (eth *EthAPI) GetCode(address common.Address, quantity string) ([]byte, error) {
	req := eth.requestManager.newRequest("eth_getCode")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return common.HexToBytes(resp.Get("result").(string)), nil
}

// Sign signs data with a given address.
func (eth *EthAPI) Sign(address common.Address, data []byte) ([]byte, error) {
	req := eth.requestManager.newRequest("eth_sign")
	req.Set("params", []string{address.String(), string(data)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return common.HexToBytes(resp.Get("result").(string)), nil
}

// SendTransaction creates new message call transaction or a contract creation,
// if the data field contains code.
func (eth *EthAPI) SendTransaction(tx *common.TransactionRequest) (hash common.Hash, err error) {
	req := eth.requestManager.newRequest("eth_sendTransaction")
	req.Set("params", []string{tx.String()})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return common.NewHash(nil), err
	}

	if resp.Error() != nil {
		return common.NewHash(nil), resp.Error()
	}

	return common.StringToHash(resp.Get("result").(string)), nil
}

// SendRawTransaction creates new message call transaction or a contract
// creation for signed transactions.
func (eth *EthAPI) SendRawTransaction(tx []byte) (hash common.Hash, err error) {
	req := eth.requestManager.newRequest("eth_sendRawTransaction")
	req.Set("params", []string{common.BytesToHex(tx)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return common.NewHash(nil), err
	}

	if resp.Error() != nil {
		return common.NewHash(nil), resp.Error()
	}

	return common.StringToHash(resp.Get("result").(string)), nil
}

// Call executes a new message call immediately without creating a transaction
// on the block chain.
func (eth *EthAPI) Call(tx *common.TransactionRequest, quantity string) ([]byte, error) {
	req := eth.requestManager.newRequest("eth_call")
	req.Set("params", []string{tx.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return common.HexToBytes(resp.Get("result").(string)), nil
}

// EstimateGas makes a call or transaction, which won't be added to the
// blockchain and returns the used gas, which can be used for estimating the
// used gas.
func (eth *EthAPI) EstimateGas(tx *common.TransactionRequest, quantity string) (result *big.Int, err error) {
	req := eth.requestManager.newRequest("eth_estimateGas")
	req.Set("params", []string{tx.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		return nil, fmt.Errorf("%v", resp.Get("result"))
	}
	return result, nil
}

// GetBlockByHash returns information about a block by hash.
func (eth *EthAPI) GetBlockByHash(hash common.Hash, full bool) (*common.Block, error) {
	req := eth.requestManager.newRequest("eth_getBlockByHash")
	req.Set("params", []interface{}{hash.String(), full})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonBlock{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToBlock(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetBlockByNumber returns information about a block by block number.
func (eth *EthAPI) GetBlockByNumber(quantity string, full bool) (*common.Block, error) {
	req := eth.requestManager.newRequest("eth_getBlockByNumber")
	req.Set("params", []interface{}{quantity, full})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonBlock{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToBlock(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetTransactionByHash returns the information about a transaction requested by
// transaction hash.
func (eth *EthAPI) GetTransactionByHash(hash common.Hash) (*common.Transaction, error) {
	req := eth.requestManager.newRequest("eth_getTransactionByHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonTransaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToTransaction(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetTransactionByBlockHashAndIndex returns information about a transaction by
// block hash and transaction index position.
func (eth *EthAPI) GetTransactionByBlockHashAndIndex(hash common.Hash, index uint64) (*common.Transaction, error) {
	req := eth.requestManager.newRequest("eth_getTransactionByBlockHashAndIndex")
	req.Set("params", []string{hash.String(), fmt.Sprintf("%v", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonTransaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToTransaction(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetTransactionByBlockNumberAndIndex returns information about a transaction
// by block number and transaction index position.
func (eth *EthAPI) GetTransactionByBlockNumberAndIndex(quantity string, index uint64) (*common.Transaction, error) {
	req := eth.requestManager.newRequest("eth_getTransactionByBlockNumberAndIndex")
	req.Set("params", []string{quantity, fmt.Sprintf("%v", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonTransaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToTransaction(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetTransactionReceipt Returns the receipt of a transaction by transaction hash.
func (eth *EthAPI) GetTransactionReceipt(hash common.Hash) (*common.TransactionReceipt, error) {
	req := eth.requestManager.newRequest("eth_getTransactionReceipt")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonTransactionReceipt{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToTransactionReceipt(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetUncleByBlockHashAndIndex returns information about a uncle of a block by
// hash and uncle index position.
func (eth *EthAPI) GetUncleByBlockHashAndIndex(hash common.Hash, index uint64) (*common.Block, error) {
	req := eth.requestManager.newRequest("eth_getUncleByBlockHashAndIndex")
	req.Set("params", []string{hash.String(), fmt.Sprintf("%d", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonBlock{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToBlock(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetUncleByBlockNumberAndIndex returns information about a uncle of a block by
// number and uncle index position.
func (eth *EthAPI) GetUncleByBlockNumberAndIndex(quantity string, index uint64) (*common.Block, error) {
	req := eth.requestManager.newRequest("eth_getUncleByBlockNumberAndIndex")
	req.Set("params", []string{quantity, fmt.Sprintf("%d", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	result := &jsonBlock{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToBlock(), nil
		}
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// GetCompilers returns a list of available compilers in the client.
func (eth *EthAPI) GetCompilers() (result []string, err error) {
	req := eth.requestManager.newRequest("eth_getCompilers")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	for _, r := range resp.Get("result").([]interface{}) {
		result = append(result, r.(string))
	}
	return result, nil
}

// NewFilter creates a filter object, based on filter options, to notify when
// the state changes (logs). To check if the state has changed, call
// eth_getFilterChanges.
func (eth *EthAPI) NewFilter(option *FilterOption) (Filter, error) {
	req := eth.requestManager.newRequest("eth_newFilter")
	if option == nil {
		option = &FilterOption{}
	}
	req.Set("params", option)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		return nil, err
	}
	return newFilter(eth, TypeNormal, id), nil
}

// NewBlockFilter creates a filter in the node, to notify when a new block
// arrives. To check if the state has changed, call eth_getFilterChanges.
func (eth *EthAPI) NewBlockFilter() (Filter, error) {
	req := eth.requestManager.newRequest("eth_newBlockFilter")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		return nil, err
	}
	return newFilter(eth, TypeBlockFilter, id), nil
}

// NewPendingTransactionFilter creates a filter in the node, to notify when new
// pending transactions arrive. To check if the state has changed, call
// eth_getFilterChanges.
func (eth *EthAPI) NewPendingTransactionFilter() (Filter, error) {
	req := eth.requestManager.newRequest("eth_newPendingTransactionFilter")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		return nil, err
	}
	return newFilter(eth, TypeTransactionFilter, id), nil
}

// UninstallFilter uninstalls a filter with given id. Should always be called
// when watch is no longer needed. Additonally Filters timeout when they aren't
// requested with eth_getFilterChanges for a period of time.
func (eth *EthAPI) UninstallFilter(filter Filter) (bool, error) {
	req := eth.requestManager.newRequest("eth_uninstallFilter")
	req.Set("params", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return false, err
	}

	if resp.Error() != nil {
		return false, resp.Error()
	}

	return resp.Get("result").(bool), nil
}

// GetFilterChanges polling method for a filter, which returns an array of logs
// which occurred since last poll.
func (eth *EthAPI) GetFilterChanges(filter Filter) (result []interface{}, err error) {
	req := eth.requestManager.newRequest("eth_getFilterChanges")
	req.Set("params", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.Get("result").([]interface{}), nil
}

// GetFilterLogs returns an array of all logs matching filter with given id.
func (eth *EthAPI) GetFilterLogs(filter Filter) (result []interface{}, err error) {
	req := eth.requestManager.newRequest("eth_getFilterLogs")
	req.Set("params", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.Get("result").([]interface{}), nil
}

// GetLogs returns an array of all logs matching a given filter object.
func (eth *EthAPI) GetLogs(filter Filter) (result []interface{}, err error) {
	req := eth.requestManager.newRequest("eth_getLogs")
	req.Set("params", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, resp.Error()
	}

	return resp.Get("result").([]interface{}), nil
}

// GetWork returns the hash of the current block, the seedHash, and the boundary
// condition to be met ("target").
func (eth *EthAPI) GetWork() (header, seed, boundary common.Hash, err error) {
	req := eth.requestManager.newRequest("eth_getWork")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return common.NewHash(nil), common.NewHash(nil), common.NewHash(nil), err
	}

	if resp.Error() != nil {
		return common.NewHash(nil), common.NewHash(nil), common.NewHash(nil), resp.Error()
	}

	results := resp.Get("result").([]interface{})
	header = common.StringToHash(results[0].(string))
	seed = common.StringToHash(results[1].(string))
	boundary = common.StringToHash(results[2].(string))
	return header, seed, boundary, nil
}

// SubmitWork is used for submitting a proof-of-work solution.
func (eth *EthAPI) SubmitWork(nonce uint64, header, mixDigest common.Hash) (bool, error) {
	req := eth.requestManager.newRequest("eth_submitWork")
	req.Set("params", []string{
		fmt.Sprintf("0x%16x", nonce),
		header.String(),
		mixDigest.String(),
	})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		return false, err
	}

	if resp.Error() != nil {
		return false, resp.Error()
	}

	return resp.Get("result").(bool), nil
}
