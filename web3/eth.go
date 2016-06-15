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

	"github.com/alanchchen/web3go/rpc"
)

// Eth ...
type Eth interface {
	ProtocolVersion() string
	Syncing() (bool, *SyncStatus)
	Coinbase() Address
	Mining() bool
	HashRate() uint64
	GasPrice() *big.Int
	Accounts() []Address
	BlockNumber() *big.Int
	GetBalance(address Address, quantity string) *big.Int
	GetStorageAt(address Address, position uint64, quantity string) uint64
	GetTransactionCount(address Address, quantity string) *big.Int
	GetBlockTransactionCountByHash(hash Hash) *big.Int
	GetBlockTransactionCountByNumber(quantity string) *big.Int
	GetUncleCountByBlockHash(hash Hash) *big.Int
	GetUncleCountByBlockNumber(quantity string) *big.Int
	GetCode(address Address, quantity string) []byte
	Sign(address Address, data []byte) []byte
	SendTransaction(tx *TransactionRequest) Hash
	SendRawTransaction(tx []byte) Hash
	Call(tx *TransactionRequest, quantity string) []byte
	EstimateGas(tx *Transaction, quantity string) *big.Int
	GetBlockByHash(hash Hash, full bool) *Block
	GetBlockByNumber(quantity string, full bool) *Block
	GetTransactionByHash(hash Hash) *Transaction
	GetTransactionByBlockHashAndIndex(hash Hash, index uint64) *Transaction
	GetTransactionByBlockNumberAndIndex(quantity string, index uint64) *Transaction
	GetTransactionReceipt(hash Hash) *TransactionReceipt
	GetUncleByBlockHashAndIndex(hash Hash, index uint64) *Block
	GetUncleByBlockNumberAndIndex(quantity string, index uint64) *Block
	GetCompilers() []string
	// GompileLLL
	// CompileSolidity
	// CompileSerpent
	NewFilter(option *FilterOption) *Filter
	NewBlockFilter() *Filter
	NewPendingTransactionFilter() *Filter
	UninstallFilter(filter *Filter) bool
	GetFilterChanges(filter *Filter) []Log
	GetFilterLogs(filter *Filter) []Log
	GetLogs(filter *Filter) []Log
	GetWork() (header Hash, seed Hash, boundary Hash)
	SubmitWork(nonce uint64, header Hash, mixDigest Hash) bool
	// SubmitHashrate
}

// EthAPI ...
type EthAPI struct {
	rpc            rpc.RPC
	requestManager *RequestManager
}

// NewEthAPI ...
func newEthAPI(requestManager *RequestManager) Eth {
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

// Syncing returns true with an object with data about the sync status or false with nil.
func (eth *EthAPI) Syncing() (bool, *SyncStatus) {
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
		r := &SyncStatus{}
		if resultBlob, err := json.Marshal(result); err == nil {
			if err := json.Unmarshal(resultBlob, &r); err == nil {
				return true, r
			}
		}
	}
	return false, nil
}

// Coinbase returns the client coinbase address.
func (eth *EthAPI) Coinbase() (addr Address) {
	req := eth.requestManager.newRequest("eth_protocolVersion")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return StringToAddress(resp.Get("result").(string))
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

// HashRate returns the number of hashes per second that the node is mining with.
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
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// Accounts returns a list of addresses owned by client.
func (eth *EthAPI) Accounts() (addrs []Address) {
	req := eth.requestManager.newRequest("eth_accounts")
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	results := resp.Get("result").([]interface{})
	for _, r := range results {
		addrs = append(addrs, StringToAddress(r.(string)))
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
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBalance returns the balance of the account of given address.
func (eth *EthAPI) GetBalance(address Address, quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_blockNumber")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetStorageAt returns the value from a storage position at a given address.
func (eth *EthAPI) GetStorageAt(address Address, position uint64, quantity string) uint64 {
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
func (eth *EthAPI) GetTransactionCount(address Address, quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getTransactionCount")
	req.Set("params", []string{address.String(), quantity})
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (eth *EthAPI) GetBlockTransactionCountByHash(hash Hash) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (eth *EthAPI) GetBlockTransactionCountByNumber(quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getBlockTransactionCountByNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (eth *EthAPI) GetUncleCountByBlockHash(hash Hash) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockHash")
	req.Set("params", hash.String())
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

// GetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (eth *EthAPI) GetUncleCountByBlockNumber(quantity string) (result *big.Int) {
	req := eth.requestManager.newRequest("eth_getUncleCountByBlockNumber")
	req.Set("params", quantity)
	resp, err := eth.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result = new(big.Int)
	_, ok := result.SetString(HexToString(resp.Get("result").(string)), 16)
	if !ok {
		panic(fmt.Errorf("Failed to parse %v", resp.Get("result")))
	}
	return result
}

func (eth *EthAPI) GetCode(address Address, quantity string) []byte {
	return nil
}

func (eth *EthAPI) Sign(address Address, data []byte) []byte {
	return nil
}

func (eth *EthAPI) SendTransaction(tx *TransactionRequest) (hash Hash) {
	hashString := "00000000000000000000000000000000"
	copy(hash[:], hashString)
	return hash
}

func (eth *EthAPI) SendRawTransaction(tx []byte) (hash Hash) {
	hashString := "00000000000000000000000000000000"
	copy(hash[:], hashString)
	return hash
}

func (eth *EthAPI) Call(tx *TransactionRequest, quantity string) []byte {
	return nil
}

func (eth *EthAPI) EstimateGas(tx *Transaction, quantity string) *big.Int {
	return nil
}

func (eth *EthAPI) GetBlockByHash(hash Hash, full bool) *Block {
	return nil
}

func (eth *EthAPI) GetBlockByNumber(quantity string, full bool) *Block {
	return nil
}

func (eth *EthAPI) GetTransactionByHash(hash Hash) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionByBlockHashAndIndex(hash Hash, index uint64) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionByBlockNumberAndIndex(quantity string, index uint64) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionReceipt(hash Hash) *TransactionReceipt {
	return nil
}

func (eth *EthAPI) GetUncleByBlockHashAndIndex(hash Hash, index uint64) *Block {
	return nil
}

func (eth *EthAPI) GetUncleByBlockNumberAndIndex(quantity string, index uint64) *Block {
	return nil
}

func (eth *EthAPI) GetCompilers() []string {
	return nil
}

func (eth *EthAPI) NewFilter(option *FilterOption) *Filter {
	return nil
}

func (eth *EthAPI) NewBlockFilter() *Filter {
	return nil
}

func (eth *EthAPI) NewPendingTransactionFilter() *Filter {
	return nil
}

func (eth *EthAPI) UninstallFilter(filter *Filter) bool {
	return false
}

func (eth *EthAPI) GetFilterChanges(filter *Filter) []Log {
	return nil
}

func (eth *EthAPI) GetFilterLogs(filter *Filter) []Log {
	return nil
}

func (eth *EthAPI) GetLogs(filter *Filter) []Log {
	return nil
}

func (eth *EthAPI) GetWork() (header Hash, seed Hash, boundary Hash) {
	hashString := "00000000000000000000000000000000"
	copy(header[:], hashString)
	copy(seed[:], hashString)
	copy(boundary[:], hashString)
	return header, seed, boundary
}

func (eth *EthAPI) SubmitWork(nonce uint64, header Hash, mixDigest Hash) bool {
	return false
}
