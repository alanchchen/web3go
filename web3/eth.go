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
	"github.com/alanchchen/web3go/filter"
	"github.com/alanchchen/web3go/rpc"
)

// Eth ...
type Eth interface {
	ProtocolVersion() string
	Syncing() (bool, *common.SyncStatus)
	Coinbase() common.Address
	Mining() bool
	HashRate() uint64
	GasPrice() *big.Int
	Accounts() []common.Address
	BlockNumber() *big.Int
	GetBalance(address common.Address, quantity string) *big.Int
	GetStorageAt(address common.Address, position uint64, quantity string) uint64
	GetTransactionCount(address common.Address, quantity string) *big.Int
	GetBlockTransactionCountByHash(hash common.Hash) *big.Int
	GetBlockTransactionCountByNumber(quantity string) *big.Int
	GetUncleCountByBlockHash(hash common.Hash) *big.Int
	GetUncleCountByBlockNumber(quantity string) *big.Int
	GetCode(address common.Address, quantity string) []byte
	Sign(address common.Address, data []byte) []byte
	SendTransaction(tx *common.TransactionRequest) common.Hash
	SendRawTransaction(tx []byte) common.Hash
	Call(tx *common.TransactionRequest, quantity string) []byte
	EstimateGas(tx *common.Transaction, quantity string) *big.Int
	GetBlockByHash(hash common.Hash, full bool) *common.Block
	GetBlockByNumber(quantity string, full bool) *common.Block
	GetTransactionByHash(hash common.Hash) *common.Transaction
	GetTransactionByBlockHashAndIndex(hash common.Hash, index uint64) *common.Transaction
	GetTransactionByBlockNumberAndIndex(quantity string, index uint64) *common.Transaction
	GetTransactionReceipt(hash common.Hash) *common.TransactionReceipt
	GetUncleByBlockHashAndIndex(hash common.Hash, index uint64) *common.Block
	GetUncleByBlockNumberAndIndex(quantity string, index uint64) *common.Block
	GetCompilers() []string
	// GompileLLL
	// CompileSolidity
	// CompileSerpent
	NewFilter(option *filter.Option) filter.Filter
	NewBlockFilter() *filter.BlockFilter
	NewPendingTransactionFilter() *filter.PendingTransactionFilter
	UninstallFilter(filter filter.Filter) bool
	GetFilterChanges(filter filter.Filter) []common.Log
	GetFilterLogs(filter filter.Filter) []common.Log
	GetLogs(filter filter.Filter) []common.Log
	GetWork() (header common.Hash, seed common.Hash, boundary common.Hash)
	SubmitWork(nonce uint64, header common.Hash, mixDigest common.Hash) bool
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
func (eth *EthAPI) ProtocolVersion() string {
	req := eth.requestManager.newRequest("eth_protocolVersion")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return resp.Get("result").(string)
}

// Syncing returns true with an object with data about the sync status or false
// with nil.
func (eth *EthAPI) Syncing() (bool, *common.SyncStatus) {
	req := eth.requestManager.newRequest("eth_syncing")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := resp.Get("result")
	switch result.(type) {
	case bool:
		return false, nil
	default:
		r := &common.SyncStatus{}
		if resultBlob, err := json.Marshal(result); err == nil {
			if err := json.Unmarshal(resultBlob, &r); err == nil {
				return true, r
			}
		}
	}
	return false, nil
}

// Coinbase returns the client coinbase address.
func (eth *EthAPI) Coinbase() (addr common.Address) {
	req := eth.requestManager.newRequest("eth_coinbase")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.StringToAddress(resp.Get("result").(string))
}

// Mining returns true if client is actively mining new blocks.
func (eth *EthAPI) Mining() bool {
	req := eth.requestManager.newRequest("eth_mining")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return resp.Get("result").(bool)
}

// HashRate returns the number of hashes per second that the node is mining
// with.
func (eth *EthAPI) HashRate() uint64 {
	req := eth.requestManager.newRequest("eth_hashrate")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result, err := strconv.ParseUint(resp.Get("result").(string), 16, 64)
	if err != nil {
		panic(err)
	}
	return result
}

// GasPrice returns the current price per gas in wei.
func (eth *EthAPI) GasPrice() (result *big.Int) {
	req := eth.requestManager.newRequest("eth_gasPrice")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// Accounts returns a list of addresses owned by client.
func (eth *EthAPI) Accounts() (addrs []common.Address) {
	req := eth.requestManager.newRequest("eth_accounts")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	results := resp.Get("result").([]interface{})
	for _, r := range results {
		addrs = append(addrs, common.StringToAddress(r.(string)))
	}
	return addrs
}

// BlockNumber returns the number of most recent block.
func (eth *EthAPI) BlockNumber() (result *big.Int) {
	req := eth.requestManager.newRequest("eth_blockNumber")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBalance returns the balance of the account of given address.
func (eth *EthAPI) GetBalance(address common.Address, quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_blockNumber")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetStorageAt returns the value from a storage position at a given address.
func (eth *EthAPI) GetStorageAt(address common.Address, position uint64, quantity string) uint64 {
	req := eth.requestManager.newRequest("eth_getStorageAt")
	req.Set("params", []string{address.String(), fmt.Sprintf("%v", position), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result, err := strconv.ParseUint(resp.Get("result").(string), 16, 64)
	if err != nil {
		panic(err)
	}
	return result
}

// GetTransactionCount returns the number of transactions sent from an address.
func (eth *EthAPI) GetTransactionCount(address common.Address, quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getTransactionCount")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBlockTransactionCountByHash returns the number of transactions in a block
// from a block matching the given block hash.
func (eth *EthAPI) GetBlockTransactionCountByHash(hash common.Hash) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBlockTransactionCountByNumber returns the number of transactions in a
// block from a block matching the given block number.
func (eth *EthAPI) GetBlockTransactionCountByNumber(quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetUncleCountByBlockHash returns the number of uncles in a block from a block
// matching the given block hash.
func (eth *EthAPI) GetUncleCountByBlockHash(hash common.Hash) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetUncleCountByBlockNumber returns the number of uncles in a block from a
// block matching the given block number.
func (eth *EthAPI) GetUncleCountByBlockNumber(quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetCode returns code at a given address.
func (eth *EthAPI) GetCode(address common.Address, quantity string) []byte {
	req := eth.requestManager.newRequest("eth_getCode")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.HexToBytes(resp.Get("result").(string))
}

// Sign signs data with a given address.
func (eth *EthAPI) Sign(address common.Address, data []byte) []byte {
	req := eth.requestManager.newRequest("eth_sign")
	req.Set("params", []string{address.String(), string(data)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.HexToBytes(resp.Get("result").(string))
}

// SendTransaction creates new message call transaction or a contract creation,
// if the data field contains code.
func (eth *EthAPI) SendTransaction(tx *common.TransactionRequest) (hash common.Hash) {
	req := eth.requestManager.newRequest("eth_sendTransaction")
	req.Set("params", []string{tx.String()})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.StringToHash(resp.Get("result").(string))
}

// SendRawTransaction creates new message call transaction or a contract
// creation for signed transactions.
func (eth *EthAPI) SendRawTransaction(tx []byte) (hash common.Hash) {
	req := eth.requestManager.newRequest("eth_sendRawTransaction")
	req.Set("params", []string{common.BytesToHex(tx)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.StringToHash(resp.Get("result").(string))
}

// Call executes a new message call immediately without creating a transaction
// on the block chain.
func (eth *EthAPI) Call(tx *common.TransactionRequest, quantity string) []byte {
	req := eth.requestManager.newRequest("eth_call")
	req.Set("params", []string{tx.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return common.HexToBytes(resp.Get("result").(string))
}

// EstimateGas makes a call or transaction, which won't be added to the
// blockchain and returns the used gas, which can be used for estimating the
// used gas.
func (eth *EthAPI) EstimateGas(tx *common.Transaction, quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_estimateGas")
	req.Set("params", []string{tx.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(common.HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBlockByHash returns information about a block by hash.
func (eth *EthAPI) GetBlockByHash(hash common.Hash, full bool) *common.Block {
	req := eth.requestManager.newRequest("eth_getBlockByHash")
	req.Set("params", []interface{}{hash.String(), full})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Block{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetBlockByNumber returns information about a block by block number.
func (eth *EthAPI) GetBlockByNumber(quantity string, full bool) *common.Block {
	req := eth.requestManager.newRequest("eth_getBlockByNumber")
	req.Set("params", []interface{}{quantity, full})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Block{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetTransactionByHash returns the information about a transaction requested by
// transaction hash.
func (eth *EthAPI) GetTransactionByHash(hash common.Hash) *common.Transaction {
	req := eth.requestManager.newRequest("eth_getTransactionByHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Transaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetTransactionByBlockHashAndIndex returns information about a transaction by
// block hash and transaction index position.
func (eth *EthAPI) GetTransactionByBlockHashAndIndex(hash common.Hash, index uint64) *common.Transaction {
	req := eth.requestManager.newRequest("eth_getTransactionByBlockHashAndIndex")
	req.Set("params", []string{hash.String(), fmt.Sprintf("%v", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Transaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetTransactionByBlockNumberAndIndex returns information about a transaction
// by block number and transaction index position.
func (eth *EthAPI) GetTransactionByBlockNumberAndIndex(quantity string, index uint64) *common.Transaction {
	req := eth.requestManager.newRequest("eth_getTransactionByBlockNumberAndIndex")
	req.Set("params", []string{quantity, fmt.Sprintf("%v", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Transaction{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetTransactionReceipt Returns the receipt of a transaction by transaction hash.
func (eth *EthAPI) GetTransactionReceipt(hash common.Hash) *common.TransactionReceipt {
	req := eth.requestManager.newRequest("eth_getTransactionReceipt")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.TransactionReceipt{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetUncleByBlockHashAndIndex returns information about a uncle of a block by
// hash and uncle index position.
func (eth *EthAPI) GetUncleByBlockHashAndIndex(hash common.Hash, index uint64) *common.Block {
	req := eth.requestManager.newRequest("eth_getUncleByBlockHashAndIndex")
	req.Set("params", []string{hash.String(), fmt.Sprintf("%d", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Block{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetUncleByBlockNumberAndIndex returns information about a uncle of a block by
// number and uncle index position.
func (eth *EthAPI) GetUncleByBlockNumberAndIndex(quantity string, index uint64) *common.Block {
	req := eth.requestManager.newRequest("eth_getUncleByBlockNumberAndIndex")
	req.Set("params", []string{quantity, fmt.Sprintf("%d", index)})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	result := &common.Block{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result
		}
	}

	panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
}

// GetCompilers returns a list of available compilers in the client.
func (eth *EthAPI) GetCompilers() (result []string) {
	req := eth.requestManager.newRequest("eth_getCompilers")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	for _, r := range resp.Get("result").([]interface{}) {
		result = append(result, r.(string))
	}
	return result
}

// NewFilter creates a filter object, based on filter options, to notify when
// the state changes (logs). To check if the state has changed, call
// eth_getFilterChanges.
func (eth *EthAPI) NewFilter(option *filter.Option) filter.Filter {
	req := eth.requestManager.newRequest("eth_newFilter")
	req.Set("params", option)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		panic(err)
	}
	return filter.NewFilter(option, id)
}

// NewBlockFilter creates a filter in the node, to notify when a new block
// arrives. To check if the state has changed, call eth_getFilterChanges.
func (eth *EthAPI) NewBlockFilter() *filter.BlockFilter {
	req := eth.requestManager.newRequest("eth_newBlockFilter")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		panic(err)
	}
	return filter.NewBlockFilter(id)
}

// NewPendingTransactionFilter creates a filter in the node, to notify when new
// pending transactions arrive. To check if the state has changed, call
// eth_getFilterChanges.
func (eth *EthAPI) NewPendingTransactionFilter() *filter.PendingTransactionFilter {
	req := eth.requestManager.newRequest("eth_newPendingTransactionFilter")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	id, err := strconv.ParseUint(common.HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		panic(err)
	}
	return filter.NewPendingTransactionFilter(id)
}

// UninstallFilter uninstalls a filter with given id. Should always be called
// when watch is no longer needed. Additonally Filters timeout when they aren't
// requested with eth_getFilterChanges for a period of time.
func (eth *EthAPI) UninstallFilter(filter filter.Filter) bool {
	req := eth.requestManager.newRequest("eth_uninstallFilter")
	req.Set("param", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	return resp.Get("result").(bool)
}

// GetFilterChanges polling method for a filter, which returns an array of logs
// which occurred since last poll.
func (eth *EthAPI) GetFilterChanges(filter filter.Filter) (result []common.Log) {
	req := eth.requestManager.newRequest("eth_getFilterChanges")
	req.Set("param", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	logs := resp.Get("result").([]interface{})
	result = make([]common.Log, len(logs))
	for i, data := range logs {
		if err := json.Unmarshal(data.([]byte), &result[i]); err != nil {
			panic(err)
		}
	}
	return result
}

// GetFilterLogs returns an array of all logs matching filter with given id.
func (eth *EthAPI) GetFilterLogs(filter filter.Filter) (result []common.Log) {
	req := eth.requestManager.newRequest("eth_getFilterLogs")
	req.Set("param", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	logs := resp.Get("result").([]interface{})
	result = make([]common.Log, len(logs))
	for i, data := range logs {
		if err := json.Unmarshal(data.([]byte), &result[i]); err != nil {
			panic(err)
		}
	}
	return result
}

// GetLogs returns an array of all logs matching a given filter object.
func (eth *EthAPI) GetLogs(filter filter.Filter) (result []common.Log) {
	req := eth.requestManager.newRequest("eth_getLogs")
	req.Set("param", fmt.Sprintf("0x%x", filter.ID()))
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	logs := resp.Get("result").([]interface{})
	result = make([]common.Log, len(logs))
	for i, data := range logs {
		if err := json.Unmarshal(data.([]byte), &result[i]); err != nil {
			panic(err)
		}
	}
	return result
}

// GetWork returns the hash of the current block, the seedHash, and the boundary
// condition to be met ("target").
func (eth *EthAPI) GetWork() (header common.Hash, seed common.Hash, boundary common.Hash) {
	req := eth.requestManager.newRequest("eth_getWork")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	results := resp.Get("result").([]string)
	header = common.StringToHash(results[0])
	seed = common.StringToHash(results[1])
	boundary = common.StringToHash(results[2])
	return header, seed, boundary
}

// SubmitWork is used for submitting a proof-of-work solution.
func (eth *EthAPI) SubmitWork(nonce uint64, header common.Hash, mixDigest common.Hash) bool {
	req := eth.requestManager.newRequest("eth_submitWork")
	req.Set("params", []string{
		fmt.Sprintf("0x%16x", nonce),
		header.String(),
		mixDigest.String(),
	})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}

	return resp.Get("result").(bool)
}
