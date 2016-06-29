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
	"testing"

	"github.com/alanchchen/web3go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type NetTestSuite struct {
	suite.Suite
	web3 *Web3
	net  Net
}

func (suite *NetTestSuite) Test_Version() {
	net := suite.net
	version, err := net.Version()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.NotEqual(suite.T(), "", version, "version is empty")
}

func (suite *NetTestSuite) Test_Listening() {
	net := suite.net
	listening, err := net.Listening()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.Exactly(suite.T(), true, listening, "should be true")
}

func (suite *NetTestSuite) Test_PeerCount() {
	net := suite.net
	count, err := net.PeerCount()
	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), 50, count, "should be equal")
}

func (suite *NetTestSuite) SetupTest() {
	suite.web3 = NewWeb3(test.NewMockHTTPProvider())
	suite.net = suite.web3.Net
}

func Test_NetTestSuite(t *testing.T) {
	suite.Run(t, new(NetTestSuite))
}
