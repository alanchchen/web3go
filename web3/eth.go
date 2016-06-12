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
	web3           *Web3
	requestManager *RequestManager
}

// NewEthAPI ...
// func NewEthAPI(web3 *Web3) Eth {
// 	return &EthAPI{web3: web3, requestManager: web3.requestManager}
// }

// func (eth *EthAPI) ProtocolVersion() string {

// }
