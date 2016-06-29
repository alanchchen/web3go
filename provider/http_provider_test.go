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

package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alanchchen/web3go/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HTTPProviderTestSuite struct {
	suite.Suite
	server   *httptest.Server
	provider Provider
}

func (suite *HTTPProviderTestSuite) Test_IsConnected() {
	provider := suite.provider
	assert.EqualValues(suite.T(), true, provider.IsConnected(), "should be equal")
}

func (suite *HTTPProviderTestSuite) Test_Send() {
	provider := suite.provider
	req := &rpc.JSONRPCRequest{
		Version:    "2.0",
		Method:     "test_method",
		Params:     nil,
		Identifier: 10}
	resp, err := provider.Send(req)

	assert.NoError(suite.T(), err, "Should be no error")
	assert.EqualValues(suite.T(), req.Version, req.Version, "should be equal")
	assert.EqualValues(suite.T(), req.Identifier, resp.ID(), "should be equal")
	assert.EqualValues(suite.T(), "ok", resp.Get("result").(string), "should be equal")
}

func (suite *HTTPProviderTestSuite) Test_GetRPCMethod() {
	provider := suite.provider
	assert.NotNil(suite.T(), provider.GetRPCMethod(), "should be equal")
}

func (suite *HTTPProviderTestSuite) SetupTest() {
	suite.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		req := rpc.JSONRPCRequest{}
		resp := rpc.JSONRPCResponse{Version: "2.0"}
		err := decoder.Decode(&req)
		if err != nil {
			resp.Identifier = 0
			resp.Result = "error"
		} else {
			resp.Identifier = req.Identifier
			switch req.Method {
			case "net_listening":
				resp.Result = true
			default:
				resp.Result = "ok"
			}
		}
		jsonBlob, _ := json.Marshal(resp)
		w.Write(jsonBlob)
	}))
	suite.provider = NewHTTPProvider(suite.server.URL, rpc.GetDefaultMethod())
}

func (suite *HTTPProviderTestSuite) TearDownTest() {
	suite.server.Close()
}

func Test_HTTPProviderTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPProviderTestSuite))
}
