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
	"fmt"

	"github.com/alanchchen/web3go/rpc"
)

// MockNetAPI ...
type MockNetAPI struct {
	rpc rpc.RPC
}

// NewMockNetAPI ...
func NewMockNetAPI(rpc rpc.RPC) MockAPI {
	return &MockNetAPI{rpc: rpc}
}

// Do ...
func (net *MockNetAPI) Do(request rpc.Request) (response rpc.Response, err error) {
	method := request.Get("method").(string)
	switch method {
	case "net_version":
		data := struct {
			Version string      `json:"version"`
			ID      uint64      `json:"id"`
			Result  interface{} `json:"result"`
		}{
			request.Get("version").(string),
			request.ID(),
			"100",
		}
		if resp := net.rpc.NewResponse(data); resp != nil {
			return resp, nil
		}
		return nil, fmt.Errorf("Failed to generate response")
	case "net_listening":
		data := struct {
			Version string      `json:"version"`
			ID      uint64      `json:"id"`
			Result  interface{} `json:"result"`
		}{
			request.Get("version").(string),
			request.ID(),
			true,
		}
		if resp := net.rpc.NewResponse(data); resp != nil {
			return resp, nil
		}
		return nil, fmt.Errorf("Failed to generate response")
	case "net_peerCount":
		data := struct {
			Version string      `json:"version"`
			ID      uint64      `json:"id"`
			Result  interface{} `json:"result"`
		}{
			request.Get("version").(string),
			request.ID(),
			"0x32",
		}
		if resp := net.rpc.NewResponse(data); resp != nil {
			return resp, nil
		}
		return nil, fmt.Errorf("Failed to generate response")
	}

	return nil, fmt.Errorf("Invalid method %s", method)
}
