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
	"math/big"

	"github.com/alanchchen/web3go/common"
)

type jsonBlock struct {
	Number          json.Number    `json:"number"`
	Hash            common.Hash    `json:"hash"`
	ParentHash      common.Hash    `json:"parentHash"`
	Nonce           common.Hash    `json:"nonce"`
	Sha3Uncles      common.Hash    `json:"sha3Uncles"`
	Bloom           common.Hash    `json:"logsBloom"`
	TransactionRoot common.Hash    `json:"transactionsRoot"`
	StateRoot       common.Hash    `json:"stateRoot"`
	Miner           common.Address `json:"miner"`
	Difficulty      json.Number    `json:"difficulty"`
	TotalDifficulty json.Number    `json:"totalDifficulty"`
	ExtraData       common.Hash    `json:"extraData"`
	Size            json.Number    `json:"size"`
	GasLimit        json.Number    `json:"gasLimit"`
	GasUsed         json.Number    `json:"gasUsed"`
	Timestamp       json.Number    `json:"timestamp"`
	Transactions    []common.Hash  `json:"transactions"`
	Uncles          []common.Hash  `json:"uncles"`
}

func (b *jsonBlock) ToBlock() (block *common.Block) {
	block = &common.Block{}
	block.Number = jsonNumbertoInt(b.Number)
	block.Hash = b.Hash
	block.ParentHash = b.ParentHash
	block.Nonce = b.Nonce
	block.Sha3Uncles = b.Sha3Uncles
	block.Bloom = b.Bloom
	block.TransactionRoot = b.TransactionRoot
	block.StateRoot = b.StateRoot
	block.Miner = b.Miner
	block.Difficulty = jsonNumbertoInt(b.Difficulty)
	block.TotalDifficulty = jsonNumbertoInt(b.TotalDifficulty)
	block.ExtraData = b.ExtraData
	block.Size = jsonNumbertoInt(b.Size)
	block.GasLimit = jsonNumbertoInt(b.GasLimit)
	block.GasUsed = jsonNumbertoInt(b.GasUsed)
	block.Timestamp = jsonNumbertoInt(b.Timestamp)
	block.Transactions = b.Transactions
	block.Uncles = b.Uncles
	return block
}

type jsonTransaction struct {
	Hash             common.Hash    `json:"hash"`
	Nonce            common.Hash    `json:"nonce"`
	BlockHash        common.Hash    `json:"blockHash"`
	BlockNumber      json.Number    `json:"blockNumber"`
	TransactionIndex uint64         `json:"transactionIndex"`
	From             common.Address `json:"from"`
	To               common.Address `json:"to"`
	Gas              json.Number    `json:"gas"`
	GasPrice         json.Number    `json:"gasprice"`
	Value            json.Number    `json:"value"`
	Data             []byte         `json:"input"`
}

func (t *jsonTransaction) ToTransaction() (tx *common.Transaction) {
	tx = &common.Transaction{}
	tx.Hash = t.Hash
	tx.Nonce = t.Nonce
	tx.BlockHash = t.BlockHash
	tx.BlockNumber = jsonNumbertoInt(t.BlockNumber)
	tx.TransactionIndex = t.TransactionIndex
	tx.From = t.From
	tx.To = t.To
	tx.Gas = jsonNumbertoInt(t.Gas)
	tx.GasPrice = jsonNumbertoInt(t.GasPrice)
	tx.Value = jsonNumbertoInt(t.Value)
	tx.Data = t.Data
	return tx
}

type jsonTransactionReceipt struct {
	Hash              common.Hash    `json:"transactionHash"`
	TransactionIndex  uint64         `json:"transactionIndex"`
	BlockNumber       json.Number    `json:"blockNumber"`
	BlockHash         common.Hash    `json:"blockHash"`
	CumulativeGasUsed json.Number    `json:"cumulativeGasUsed"`
	GasUsed           json.Number    `json:"gasUsed"`
	ContractAddress   common.Address `json:"contractAddress"`
	Logs              []jsonLog      `json:"logs"`
}

func (r *jsonTransactionReceipt) ToTransactionReceipt() (receipt *common.TransactionReceipt) {
	receipt = &common.TransactionReceipt{}
	receipt.Hash = r.Hash
	receipt.TransactionIndex = r.TransactionIndex
	receipt.BlockNumber = jsonNumbertoInt(r.BlockNumber)
	receipt.BlockHash = r.BlockHash
	receipt.CumulativeGasUsed = jsonNumbertoInt(r.CumulativeGasUsed)
	receipt.GasUsed = jsonNumbertoInt(r.GasUsed)
	receipt.ContractAddress = r.ContractAddress
	receipt.Logs = make([]common.Log, 0)
	for _, l := range r.Logs {
		receipt.Logs = append(receipt.Logs, l.ToLog())
	}
	return receipt
}

type jsonLog struct {
	LogIndex         uint64         `json:"logIndex"`
	BlockNumber      json.Number    `json:"blockNumber"`
	BlockHash        common.Hash    `json:"blockHash"`
	TransactionHash  common.Hash    `json:"transactionHash"`
	TransactionIndex uint64         `json:"transactionIndex"`
	Address          common.Address `json:"address"`
	Data             []byte         `json:"data"`
	Topics           common.Topics  `json:"topics"`
}

func (l jsonLog) ToLog() (log common.Log) {
	log = common.Log{}
	log.LogIndex = l.LogIndex
	log.BlockNumber = jsonNumbertoInt(l.BlockNumber)
	log.BlockHash = l.BlockHash
	log.TransactionHash = l.TransactionHash
	log.TransactionIndex = l.TransactionIndex
	log.Address = l.Address
	log.Data = l.Data
	log.Topics = l.Topics
	return log
}

func jsonNumbertoInt(data json.Number) *big.Int {
	f := big.NewFloat(0.0)
	f.SetString(string(data))
	result, _ := f.Int(nil)
	return result
}
