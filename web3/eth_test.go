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
	"testing"

	"github.com/alanchchen/web3go/common"
	"github.com/alanchchen/web3go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EthTestSuite struct {
	suite.Suite
	web3 *Web3
	eth  Eth
}

func (suite *EthTestSuite) Test_ProcotolVersion() {
	eth := suite.eth
	assert.NotEqual(suite.T(), "", eth.ProtocolVersion(), "version is empty")
}

func (suite *EthTestSuite) Test_Syncing() {
	eth := suite.eth
	ok, _ := eth.Syncing()
	assert.Exactly(suite.T(), false, ok, "should be false")
}

func (suite *EthTestSuite) Test_Coinbase() {
	eth := suite.eth
	address := eth.Coinbase()
	assert.EqualValues(suite.T(), "0x407d73d8a49eeb85d32cf465507dd71d507100c1", address.String(), "should be equal")
}

func (suite *EthTestSuite) Test_Mining() {
	eth := suite.eth
	assert.EqualValues(suite.T(), true, eth.Mining(), "should be equal")
}

func (suite *EthTestSuite) Test_HashRate() {
	eth := suite.eth
	assert.EqualValues(suite.T(), 0x38a, eth.HashRate(), "Should be equal")
}

func (suite *EthTestSuite) Test_GasPrice() {
	eth := suite.eth
	assert.EqualValues(suite.T(), big.NewInt(0x09184e72a000), eth.GasPrice(), "Should be equal")
}

func (suite *EthTestSuite) Test_Accounts() {
	eth := suite.eth
	assert.EqualValues(suite.T(), []common.Address{
		common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
		common.NewAddress(common.HexToBytes("0x407d73d8a49ee783afd32cf465507dd71d507100")),
	}, eth.Accounts(), "Should be equal")
}

func (suite *EthTestSuite) Test_BlockNumber() {
	eth := suite.eth
	assert.EqualValues(suite.T(), big.NewInt(0x4b7), eth.BlockNumber(), "Should be equal")
}

func (suite *EthTestSuite) Test_GetBalance() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0x0234c8a3397aab58),
		eth.GetBalance(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), "latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetStorageAt() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		0x03,
		eth.GetStorageAt(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), 0, "latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetTransactionCount() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		eth.GetTransactionCount(common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")), "latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockTransactionCountByHash() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0xb),
		eth.GetBlockTransactionCountByHash(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"))),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockTransactionCountByNumber() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0xa),
		eth.GetBlockTransactionCountByNumber("latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleCountByBlockHash() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		eth.GetUncleCountByBlockHash(common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"))),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetUncleCountByBlockNumber() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		big.NewInt(0x1),
		eth.GetUncleCountByBlockNumber("latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetCode() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x600160008035811a818181146012578301005b601b6001356025565b8060005260206000f25b600060078202905091905056"),
		eth.GetCode(common.NewAddress(common.HexToBytes("0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b")), "0x2"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_Sign() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x2ac19db245478a06032e69cdbd2b54e648b78431d0a47bd1fbab18f79f820ba407466e37adbe9e84541cab97ab7d290f4a64a5825c876d22109f3bf813254e8601"),
		eth.Sign(common.NewAddress(common.HexToBytes("0xd1ade25ccd3d550a7eb532ac759cac7be09c2719")), []byte("Schoolbus")),
		"Should be equal")
}

func (suite *EthTestSuite) Test_SendTransaction() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	assert.EqualValues(suite.T(),
		common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		eth.SendTransaction(req),
		"Should be equal")
}

func (suite *EthTestSuite) Test_SendRawTransaction() {
	eth := suite.eth
	assert.EqualValues(suite.T(),
		common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		eth.SendRawTransaction(common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675")),
		"Should be equal")
}

func (suite *EthTestSuite) Test_Call() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	assert.EqualValues(suite.T(),
		common.HexToBytes("0x"),
		eth.Call(req, "latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_EstimateGas() {
	eth := suite.eth
	req := &common.TransactionRequest{
		From:     common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
		To:       common.NewAddress(common.HexToBytes("0xd46e8dd67c5d32be8058bb8eb970870f07244567")),
		Gas:      big.NewInt(0x76c0),
		GasPrice: big.NewInt(0x9184e72a000),
		Value:    big.NewInt(0x9184e72a),
		Data:     common.HexToBytes("0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"),
	}
	assert.EqualValues(suite.T(),
		big.NewInt(0x5208),
		eth.EstimateGas(req, "latest"),
		"Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockByHash() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock := eth.GetBlockByHash(common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")), true)

	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) Test_GetBlockByNumber() {
	eth := suite.eth
	block := &common.Block{
		Number:          big.NewInt(0x1b4),
		Hash:            common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		ParentHash:      common.NewHash(common.HexToBytes("0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5")),
		Nonce:           common.NewHash(common.HexToBytes("0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2")),
		Sha3Uncles:      common.NewHash(common.HexToBytes("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")),
		Bloom:           common.NewHash(common.HexToBytes("0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")),
		TransactionRoot: common.NewHash(common.HexToBytes("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")),
		StateRoot:       common.NewHash(common.HexToBytes("0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff")),
		Miner:           common.NewAddress(common.HexToBytes("0x4e65fda2159562a496f9f3522f89122a3088497a")),
		Difficulty:      big.NewInt(0x027f07),
		TotalDifficulty: big.NewInt(0x027f07),
		ExtraData:       common.NewHash(common.HexToBytes("0x0000000000000000000000000000000000000000000000000000000000000000")),
		Size:            big.NewInt(0x027f07),
		GasLimit:        big.NewInt(0x9f759),
		GasUsed:         big.NewInt(0x9f759),
		Timestamp:       big.NewInt(0x54e34e8e),
		Transactions:    []common.Hash{},
		Uncles:          []common.Hash{},
	}
	returnedBlock := eth.GetBlockByNumber("0x1b4", true)

	assert.EqualValues(suite.T(),
		block, returnedBlock, "Should be equal")
}

func (suite *EthTestSuite) SetupTest() {
	suite.web3 = NewWeb3(test.NewMockHTTPProvider())
	suite.eth = suite.web3.Eth
}

func Test_EthTestSuite(t *testing.T) {
	suite.Run(t, new(EthTestSuite))
}
