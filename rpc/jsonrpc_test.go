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

package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JSONRPCTestSuite struct {
	suite.Suite
	rpc RPC
}

func (suite *JSONRPCTestSuite) Test_Name() {
	rpc := suite.rpc
	assert.EqualValues(suite.T(), "jsonrpc", rpc.Name(), "Should be equal")
}

func (suite *JSONRPCTestSuite) Test_NewRequest() {
	rpc := suite.rpc
	req := rpc.NewRequest("test", "arg1", "arg2")

	if assert.NotNil(suite.T(), req) {
		assert.EqualValues(suite.T(), "2.0", req.Get("version").(string), "Should be equal")
		assert.EqualValues(suite.T(), []interface{}{"arg1", "arg2"}, req.Get("params").([]interface{}), "Should be equal")
		assert.EqualValues(suite.T(), "test", req.Get("method").(string), "Should be equal")
		assert.NotZero(suite.T(), req.ID(), "Should not be zero")
		assert.EqualValues(suite.T(), `{"jsonrpc":"2.0","method":"test","params":["arg1","arg2"],"id":1}`, req.String(), "Should be equal")

		req.Set("method", "new_method")
		assert.EqualValues(suite.T(), "new_method", req.Get("method").(string), "Should be equal")
		req.Set("params", "test_params")
		assert.EqualValues(suite.T(), []interface{}{"test_params"}, req.Get("params").([]interface{}), "Should be equal")
		req.Set("params", []string{"test_param1", "test_param2"})
		assert.EqualValues(suite.T(), []interface{}{"test_param1", "test_param2"}, req.Get("params").([]interface{}), "Should be equal")
	}
}

func (suite *JSONRPCTestSuite) Test_NewResponse() {
	rpc := suite.rpc
	resp := rpc.NewResponse([]byte(`{"jsonrpc": "2.0", "id": 1, "result": ["result1", "result2"]}`))
	if assert.NotNil(suite.T(), resp) {
		assert.EqualValues(suite.T(), "2.0", resp.Get("version").(string), "Should be equal")
		assert.EqualValues(suite.T(), 1, resp.ID(), "Should be equal")

		var results []string
		for _, r := range resp.Get("result").([]interface{}) {
			results = append(results, r.(string))
		}
		assert.EqualValues(suite.T(), []string{"result1", "result2"}, results, "Should be equal")
		assert.EqualValues(suite.T(), `{"jsonrpc":"2.0","id":1,"result":["result1","result2"]}`, resp.String(), "Should be equal")
	}

	resp = rpc.NewResponse([]byte(`{"jsonrpc": "2.0", "id": 1, "result": ["result1", "result2"]}`))
	if assert.NotNil(suite.T(), resp) {
		assert.EqualValues(suite.T(), "2.0", resp.Get("version").(string), "Should be equal")
		assert.EqualValues(suite.T(), 1, resp.ID(), "Should be equal")

		var results []string
		for _, r := range resp.Get("result").([]interface{}) {
			results = append(results, r.(string))
		}
		assert.EqualValues(suite.T(), []string{"result1", "result2"}, results, "Should be equal")
		assert.Nil(suite.T(), resp.Get("xxx"))
		assert.EqualValues(suite.T(), `{"jsonrpc":"2.0","id":1,"result":["result1","result2"]}`, resp.String(), "Should be equal")
	}

	resp = rpc.NewResponse([]byte("xxx"))
	assert.Nil(suite.T(), resp)
}

func (suite *JSONRPCTestSuite) SetupTest() {
	suite.rpc = NewJSONRPC()
}

func (suite *JSONRPCTestSuite) TearDownTest() {

}

func Test_JSONRPCTestSuite(t *testing.T) {
	suite.Run(t, new(JSONRPCTestSuite))
}
