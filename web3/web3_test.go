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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Web3TestSuite struct {
	suite.Suite
	web3 *Web3
}

func (suite *Web3TestSuite) Test_IsConnected() {
	web3 := suite.web3
	assert.Equal(suite.T(), web3.IsConnected(), true, "should be true")
}

func (suite *Web3TestSuite) Test_Sha3() {
	web3 := suite.web3
	s := "Some string to be hashed"
	hash := web3.Sha3(s, nil)
	assert.Equal(suite.T(), "0xed973b234cf2238052c9ac87072c71bcf33abc1bbd721018e0cca448ef79b379", hash, "should be equal")

	encoding := struct {
		Encoding string `json:"encoding"`
	}{
		Encoding: "hex",
	}
	assert.Equal(suite.T(), "0x85dd39c91a64167ba20732b228251e67caed1462d4bcf036af88dc6856d0fdcc", web3.Sha3(hash, encoding), "should be equal")
}

func (suite *Web3TestSuite) Test_ToHex() {
	web3 := suite.web3
	s := web3.ToHex("{\"test\":\"test\"}")
	assert.Equal(suite.T(), "0x7b2274657374223a2274657374227d", s, "should be equal")

	d := struct {
		Test string `json:"test"`
	}{
		Test: "test",
	}
	s = web3.ToHex(d)
	assert.Equal(suite.T(), "0x7b2274657374223a2274657374227d", s, "should be equal")

	s = web3.ToHex(true)
	assert.Equal(suite.T(), "0x1", s, "should be equal")

	s = web3.ToHex(false)
	assert.Equal(suite.T(), "0x0", s, "should be equal")

	i := new(big.Int)
	i.SetInt64(12345)
	s = web3.ToHex(i)
	assert.Equal(suite.T(), "0x3039", s, "should be equal")
}

func (suite *Web3TestSuite) Test_ToASCII() {
	web3 := suite.web3
	s := "0x657468657265756d000000000000000000000000000000000000000000000000"
	assert.Equal(suite.T(), "ethereum", web3.ToASCII(s), "should be equal")
}

func (suite *Web3TestSuite) Test_FromASCII() {
	web3 := suite.web3
	s := "ethereum"
	assert.Equal(suite.T(), "0x657468657265756d", web3.FromASCII(s, 0), "should be equal")
	assert.Equal(suite.T(), "0x657468657265756d000000000000000000000000000000000000000000000000", web3.FromASCII(s, 32), "should be equal")
}

func (suite *Web3TestSuite) Test_ToDecimal() {
	web3 := suite.web3
	s := "0x15"
	assert.Equal(suite.T(), "21", web3.ToDecimal(s), "should be equal")
}

func (suite *Web3TestSuite) Test_FromDecimal() {
	web3 := suite.web3
	s := "21"
	assert.Equal(suite.T(), "0x15", web3.FromDecimal(s), "should be equal")
}

func (suite *Web3TestSuite) Test_FromWei() {
	web3 := suite.web3
	s := "1"
	assert.Equal(suite.T(), "1", web3.FromWei(s, "wei"), "should be equal")
	s = "1000000"
	assert.Equal(suite.T(), "1000", web3.FromWei(s, "kwei"), "should be equal")
	s = "999000000000"
	assert.Equal(suite.T(), "999", web3.FromWei(s, "gwei"), "should be equal")
	s = "123000000000000000000"
	assert.Equal(suite.T(), "123", web3.FromWei(s, "ether"), "should be equal")
	s = "1000000000000000000000000000000"
	assert.Equal(suite.T(), "1", web3.FromWei(s, "tether"), "should be equal")
}

func (suite *Web3TestSuite) Test_ToWei() {
	web3 := suite.web3
	s := "100000"
	assert.Equal(suite.T(), "0", web3.ToWei(s, "noether"), "should be equal")
	s = "1"
	assert.Equal(suite.T(), "1", web3.ToWei(s, "wei"), "should be equal")
	s = "1000"
	assert.Equal(suite.T(), "1000000", web3.ToWei(s, "kwei"), "should be equal")
	s = "999"
	assert.Equal(suite.T(), "999000000000", web3.ToWei(s, "gwei"), "should be equal")
	s = "123"
	assert.Equal(suite.T(), "123000000000000000000", web3.ToWei(s, "ether"), "should be equal")
	s = "1"
	assert.Equal(suite.T(), "1000000000000000000000000000000", web3.ToWei(s, "tether"), "should be equal")
}

func (suite *Web3TestSuite) Test_ToBigNumber() {
	web3 := suite.web3
	s := "200000000000000000000001"
	f64s, _ := web3.ToBigNumber(s).Float64()
	assert.Equal(suite.T(), 2.0000000000000002e+23, f64s, "should be equal")
}

func (suite *Web3TestSuite) Test_IsAddress() {
	web3 := suite.web3
	s := "0x1a9afb711302c5f83b5902843d1c007a1a137632"
	assert.Equal(suite.T(), true, web3.IsAddress(s), "should be equal")
	s = "0x26c7ea56af25113f712befbf2077798fd7fbdb7c"
	assert.Equal(suite.T(), true, web3.IsAddress(s), "should be equal")
	s = "0xa4137d4ad166ae825f1b8dbb0c3d48f25f172e9e"
	assert.Equal(suite.T(), true, web3.IsAddress(s), "should be equal")
	s = "0xA4137D4AD166AE825F1B8DBB0C3D48F25F172E9E"
	assert.Equal(suite.T(), true, web3.IsAddress(s), "should be equal")
	s = "0x1"
	assert.Equal(suite.T(), false, web3.IsAddress(s), "should be equal")
	s = "0xA4137d4ad166ae825f1b8dbb0c3d48f25f172e9e"
	assert.Equal(suite.T(), false, web3.IsAddress(s), "should be equal")
	s = "I'm an account"
	assert.Equal(suite.T(), false, web3.IsAddress(s), "should be equal")
}

func (suite *Web3TestSuite) SetupTest() {
	suite.web3 = NewWeb3(NewMockHTTPProvider())
}

func Test_Web3TestSuite(t *testing.T) {
	suite.Run(t, new(Web3TestSuite))
}
