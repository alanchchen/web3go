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

package test

import (
	"fmt"
	"math/big"

	"github.com/alanchchen/web3go/common"
	"github.com/alanchchen/web3go/rpc"
)

// MockEthAPI ...
type MockEthAPI struct {
	rpc rpc.RPC
}

// NewMockEthAPI ...
func NewMockEthAPI(rpc rpc.RPC) MockAPI {
	return &MockEthAPI{rpc: rpc}
}

// Do ...
func (eth *MockEthAPI) Do(request rpc.Request) (response rpc.Response, err error) {
	method := request.Get("method").(string)
	switch method {
	case "eth_protocolVersion":
		return generateResponse(eth.rpc, request, "54")
	case "eth_syncing":
		return generateResponse(eth.rpc, request, false)
	case "eth_coinbase":
		return generateResponse(eth.rpc, request, "0x407d73d8a49eeb85d32cf465507dd71d507100c1")
	case "eth_mining":
		return generateResponse(eth.rpc, request, true)
	case "eth_hashrate":
		return generateResponse(eth.rpc, request, "0x38a")
	case "eth_gasPrice":
		return generateResponse(eth.rpc, request, "0x09184e72a000")
	case "eth_accounts":
		return generateResponse(eth.rpc, request,
			[]string{"0x407d73d8a49eeb85d32cf465507dd71d507100c1",
				"0x407d73d8a49ee783afd32cf465507dd71d507100"})
	case "eth_blockNumber":
		return generateResponse(eth.rpc, request, "0x4b7")
	case "eth_getBalance":
		return generateResponse(eth.rpc, request, "0x0234c8a3397aab58")
	case "eth_getStorageAt":
		return generateResponse(eth.rpc, request, "0x03")
	case "eth_getTransactionCount":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_getBlockTransactionCountByHash":
		return generateResponse(eth.rpc, request, "0xb")
	case "eth_getBlockTransactionCountByNumber":
		return generateResponse(eth.rpc, request, "0xa")
	case "eth_getUncleCountByBlockHash":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_getUncleCountByBlockNumber":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_getCode":
		return generateResponse(eth.rpc, request, "0x600160008035811a818181146012578301005b601b6001356025565b8060005260206000f25b600060078202905091905056")
	case "eth_sign":
		return generateResponse(eth.rpc, request, "0x2ac19db245478a06032e69cdbd2b54e648b78431d0a47bd1fbab18f79f820ba407466e37adbe9e84541cab97ab7d290f4a64a5825c876d22109f3bf813254e8601")
	case "eth_sendTransaction":
		return generateResponse(eth.rpc, request, "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")
	case "eth_sendRawTransaction":
		return generateResponse(eth.rpc, request, "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331")
	case "eth_call":
		return generateResponse(eth.rpc, request, "0x")
	case "eth_estimateGas":
		return generateResponse(eth.rpc, request, "0x5208")
	case "eth_getBlockByHash":
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
		return generateResponse(eth.rpc, request, block)
	case "eth_getBlockByNumber":
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
		return generateResponse(eth.rpc, request, block)
	case "eth_getTransactionByHash":
		tx := &common.Transaction{
			Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
			Nonce:            common.NewHash(common.HexToBytes("0x")),
			BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
			BlockNumber:      big.NewInt(0x15df),
			TransactionIndex: 0x1,
			From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
			To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
			Value:            big.NewInt(0x7f110),
			Gas:              big.NewInt(0x7f110),
			GasPrice:         big.NewInt(0x09184e72a000),
			Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
		}
		return generateResponse(eth.rpc, request, tx)
	case "eth_getTransactionByBlockHashAndIndex":
		tx := &common.Transaction{
			Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
			Nonce:            common.NewHash(common.HexToBytes("0x")),
			BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
			BlockNumber:      big.NewInt(0x15df),
			TransactionIndex: 0x1,
			From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
			To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
			Value:            big.NewInt(0x7f110),
			Gas:              big.NewInt(0x7f110),
			GasPrice:         big.NewInt(0x09184e72a000),
			Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
		}
		return generateResponse(eth.rpc, request, tx)
	case "eth_getTransactionByBlockNumberAndIndex":
		tx := &common.Transaction{
			Hash:             common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
			Nonce:            common.NewHash(common.HexToBytes("0x")),
			BlockHash:        common.NewHash(common.HexToBytes("0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b")),
			BlockNumber:      big.NewInt(0x15df),
			TransactionIndex: 0x1,
			From:             common.NewAddress(common.HexToBytes("0x407d73d8a49eeb85d32cf465507dd71d507100c1")),
			To:               common.NewAddress(common.HexToBytes("0x85h43d8a49eeb85d32cf465507dd71d507100c1")),
			Value:            big.NewInt(0x7f110),
			Gas:              big.NewInt(0x7f110),
			GasPrice:         big.NewInt(0x09184e72a000),
			Data:             common.HexToBytes("0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360"),
		}
		return generateResponse(eth.rpc, request, tx)
	case "eth_getTransactionReceipt":
		receipt := &common.TransactionReceipt{
			Hash:              common.NewHash(common.HexToBytes("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")),
			TransactionIndex:  0x1,
			BlockNumber:       big.NewInt(0xb),
			BlockHash:         common.NewHash(common.HexToBytes("0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b")),
			CumulativeGasUsed: big.NewInt(0x33bc),
			GasUsed:           big.NewInt(0x4dc),
			ContractAddress:   common.NewAddress(common.HexToBytes("0xb60e8dd61c5d32be8058bb8eb970870f07233155")),
			Logs:              []common.Log{},
		}
		return generateResponse(eth.rpc, request, receipt)
	case "eth_getUncleByBlockHashAndIndex":
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
		return generateResponse(eth.rpc, request, block)
	case "eth_getUncleByBlockNumberAndIndex":
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
		return generateResponse(eth.rpc, request, block)
	case "eth_getCompilers":
		return generateResponse(eth.rpc, request, []string{"solidity", "lll", "serpent"})
	// case "eth_compileSolidity":
	// case "eth_compileLLL":
	// case "eth_compileSerpent":
	case "eth_newFilter":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_newBlockFilter":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_newPendingTransactionFilter":
		return generateResponse(eth.rpc, request, "0x1")
	case "eth_uninstallFilter":
		return generateResponse(eth.rpc, request, true)
	case "eth_getFilterChanges":
		logs := []common.Log{
			{
				LogIndex:         0x1,
				BlockNumber:      big.NewInt(0x1b4),
				BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
				TransactionIndex: 0,
				Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
				Topics: common.Topics{
					{
						Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
					},
				},
			},
		}
		return generateResponse(eth.rpc, request, logs)
	case "eth_getFilterLogs":
		logs := []common.Log{
			{
				LogIndex:         0x1,
				BlockNumber:      big.NewInt(0x1b4),
				BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
				TransactionIndex: 0,
				Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
				Topics: common.Topics{
					{
						Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
					},
				},
			},
		}
		return generateResponse(eth.rpc, request, logs)
	case "eth_getLogs":
		logs := []common.Log{
			{
				LogIndex:         0x1,
				BlockNumber:      big.NewInt(0x1b4),
				BlockHash:        common.NewHash(common.HexToBytes("0x8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				TransactionHash:  common.NewHash(common.HexToBytes("0xdf829c5a142f1fccd7d8216c5785ac562ff41e2dcfdf5785ac562ff41e2dcf")),
				TransactionIndex: 0,
				Address:          common.NewAddress(common.HexToBytes("0x16c5785ac562ff41e2dcfdf829c5a142f1fccd7d")),
				Data:             []byte("0000000000000000000000000000000000000000000000000000000000000000"),
				Topics: common.Topics{
					{
						Data: common.HexToBytes("0x59ebeb90bc63057b6515673c3ecf9438e5058bca0f92585014eced636878c9a5"),
					},
				},
			},
		}
		return generateResponse(eth.rpc, request, logs)
	case "eth_getWork":
		return generateResponse(eth.rpc, request, []string{
			"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			"0x5EED00000000000000000000000000005EED0000000000000000000000000000",
			"0xd1ff1c01710000000000000000000000d1ff1c01710000000000000000000000"})
	case "eth_submitWork":
		return generateResponse(eth.rpc, request, true)

		// case "eth_submitHashrate":
	}

	return nil, fmt.Errorf("Invalid method %s", method)
}
