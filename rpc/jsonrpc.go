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
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync/atomic"
)

var (
	big1    = big.NewInt(1)
	version = "2.0"
)

// JSONRPCRequest ...
type JSONRPCRequest struct {
	Version    string        `json:"jsonrpc,omitempty"`
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	Identifier uint64        `json:"id"`
}

// Set ...
func (req *JSONRPCRequest) Set(key string, value interface{}) {
	k := strings.ToLower(key)
	switch k {
	case "method":
		req.Method = fmt.Sprintf("%v", value)
	case "params":
		req.Params = req.Params[:0]
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice, reflect.Array:
			v := reflect.ValueOf(value)
			for i := 0; i < v.Len(); i++ {
				req.Params = append(req.Params, v.Index(i).String())
			}
		default:
			req.Params = append(req.Params, value)
		}
	}
}

// Get ...
func (req *JSONRPCRequest) Get(key string) interface{} {
	k := strings.ToLower(key)
	switch k {
	case "version":
		return req.Version
	case "method":
		return req.Method
	case "params":
		return req.Params
	case "id":
		return req.Identifier
	}

	return nil
}

// String ...
func (req *JSONRPCRequest) String() string {
	jsonBytes, _ := json.Marshal(req)
	return string(jsonBytes)
}

// ID ...
func (req *JSONRPCRequest) ID() uint64 {
	return req.Identifier
}

// -----------------------------------------------------------------------------

// JSONRPCError ...
type JSONRPCError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (err *JSONRPCError) Error() string {
	return err.Message
}

// JSONRPCResponse ...
type JSONRPCResponse struct {
	Version    string        `json:"jsonrpc"`
	Identifier uint64        `json:"id"`
	Result     interface{}   `json:"result"`
	Err        *JSONRPCError `json:"error,omitempty"`
}

// Get ...
func (resp *JSONRPCResponse) Get(key string) interface{} {
	k := strings.ToLower(key)
	switch k {
	case "version":
		return resp.Version
	case "id":
		return resp.Identifier
	case "result":
		return resp.Result
	case "error":
		return resp.Err
	}

	return nil
}

// String ...
func (resp *JSONRPCResponse) String() string {
	jsonBytes, _ := json.Marshal(resp)
	return string(jsonBytes)
}

// ID ...
func (resp *JSONRPCResponse) ID() uint64 {
	return resp.Identifier
}

func (resp *JSONRPCResponse) Error() error {
	if resp.Err != nil && resp.Err.Code != 0 {
		return resp.Err
	}
	return nil
}

// -----------------------------------------------------------------------------

// JSONRPC ...
type JSONRPC struct {
	messageID uint64
}

// NewJSONRPC ...
func NewJSONRPC() RPC {
	return &JSONRPC{messageID: 0}
}

// Name ...
func (rpc *JSONRPC) Name() string {
	return "jsonrpc"
}

// NewRequest ...
func (rpc *JSONRPC) NewRequest(method string, args ...interface{}) Request {
	request := &JSONRPCRequest{Version: version, Method: method, Identifier: rpc.newID()}
	request.Params = make([]interface{}, 0)
	for _, arg := range args {
		request.Params = append(request.Params, fmt.Sprintf("%v", arg))
	}
	return request
}

// NewResponse ...
func (rpc *JSONRPC) NewResponse(data []byte) Response {
	resp := &JSONRPCResponse{}
	if err := json.Unmarshal(data, &resp); err == nil {
		return resp
	}

	return nil
}

func (rpc *JSONRPC) newID() uint64 {
	return atomic.AddUint64(&rpc.messageID, 1)
}
