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
	"math/big"

	"github.com/alanchchen/web3go/rpc"
)

// Eth ...
type Eth interface {
	ProtocolVersion() string
	Syncing() bool
	Coinbase() Address
	Mining() bool
	HashRate() uint64
	GasPrice() *big.Int
	Accounts() []Address
	BlockNumber() *big.Int
	GetBalance(address Address, quantity *BlockIndicator) *big.Int
	GetStorageAt(address Address, position uint64, quantity *BlockIndicator) []byte
	GetTransactionCount(address Address, quantity *BlockIndicator) *big.Int
	GetBlockTransactionCountByHash(hash Hash) *big.Int
	GetBlockTransactionCountByNumber(quantity *BlockIndicator) *big.Int
	GetUncleCountByBlockHash(hash Hash) *big.Int
	GetUncleCountByBlockNumber(quantity *BlockIndicator) *big.Int
	GetCode(address Address, quantity *BlockIndicator) []byte
	Sign(address Address, data []byte) []byte
	SendTransaction(tx *TransactionRequest) Hash
	SendRawTransaction(tx []byte) Hash
	Call(tx *TransactionRequest, quantity *BlockIndicator) []byte
	EstimateGas(tx *Transaction, quantity *BlockIndicator) *big.Int
	GetBlockByHash(hash Hash, full bool) *Block
	GetBlockByNumber(quantity *BlockIndicator, full bool) *Block
	GetTransactionByHash(hash Hash) *Transaction
	GetTransactionByBlockHashAndIndex(hash Hash, index uint64) *Transaction
	GetTransactionByBlockNumberAndIndex(quantity *BlockIndicator, index uint64) *Transaction
	GetTransactionReceipt(hash Hash) *TransactionReceipt
	GetUncleByBlockHashAndIndex(hash Hash, index uint64) *Block
	GetUncleByBlockNumberAndIndex(quantity *BlockIndicator, index uint64) *Block
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

func (eth *EthAPI) ProtocolVersion() string {
	return ""
}

func (eth *EthAPI) Syncing() bool {
	return false
}

func (eth *EthAPI) Coinbase() (addr Address) {
	addrString := "00000000000000000000"
	copy(addr[:], addrString)
	return addr
}

func (eth *EthAPI) Mining() bool {
	return false
}

func (eth *EthAPI) HashRate() uint64 {
	return 0
}

func (eth *EthAPI) GasPrice() *big.Int {
	return nil
}

func (eth *EthAPI) Accounts() []Address {
	return nil
}

func (eth *EthAPI) BlockNumber() *big.Int {
	return nil
}

func (eth *EthAPI) GetBalance(address Address, quantity *BlockIndicator) *big.Int {
	return nil
}

func (eth *EthAPI) GetStorageAt(address Address, position uint64, quantity *BlockIndicator) []byte {
	return nil
}

func (eth *EthAPI) GetTransactionCount(address Address, quantity *BlockIndicator) *big.Int {
	return nil
}

func (eth *EthAPI) GetBlockTransactionCountByHash(hash Hash) *big.Int {
	return nil
}

func (eth *EthAPI) GetBlockTransactionCountByNumber(quantity *BlockIndicator) *big.Int {
	return nil
}

func (eth *EthAPI) GetUncleCountByBlockHash(hash Hash) *big.Int {
	return nil
}

func (eth *EthAPI) GetUncleCountByBlockNumber(quantity *BlockIndicator) *big.Int {
	return nil
}

func (eth *EthAPI) GetCode(address Address, quantity *BlockIndicator) []byte {
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

func (eth *EthAPI) Call(tx *TransactionRequest, quantity *BlockIndicator) []byte {
	return nil
}

func (eth *EthAPI) EstimateGas(tx *Transaction, quantity *BlockIndicator) *big.Int {
	return nil
}

func (eth *EthAPI) GetBlockByHash(hash Hash, full bool) *Block {
	return nil
}

func (eth *EthAPI) GetBlockByNumber(quantity *BlockIndicator, full bool) *Block {
	return nil
}

func (eth *EthAPI) GetTransactionByHash(hash Hash) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionByBlockHashAndIndex(hash Hash, index uint64) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionByBlockNumberAndIndex(quantity *BlockIndicator, index uint64) *Transaction {
	return nil
}

func (eth *EthAPI) GetTransactionReceipt(hash Hash) *TransactionReceipt {
	return nil
}

func (eth *EthAPI) GetUncleByBlockHashAndIndex(hash Hash, index uint64) *Block {
	return nil
}

func (eth *EthAPI) GetUncleByBlockNumberAndIndex(quantity *BlockIndicator, index uint64) *Block {
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
