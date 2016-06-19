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

import "strconv"

// Net ...
type Net interface {
	Version() string
	PeerCount() uint64
	Listening() bool
}

// NetAPI ...
type NetAPI struct {
	requestManager *requestManager
}

// NewNetAPI ...
func newNetAPI(requestManager *requestManager) Net {
	return &NetAPI{requestManager: requestManager}
}

// Version returns the current network protocol version.
func (net *NetAPI) Version() string {
	req := net.requestManager.newRequest("net_version")
	resp, err := net.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return resp.Get("result").(string)
}

// PeerCount returns number of peers currenly connected to the client.
func (net *NetAPI) PeerCount() uint64 {
	req := net.requestManager.newRequest("net_peerCount")
	resp, err := net.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	result, err := strconv.ParseUint(HexToString(resp.Get("result").(string)), 16, 64)
	if err != nil {
		panic(err)
	}
	return result
}

// Listening returns true if client is actively listening for network connections.
func (net *NetAPI) Listening() bool {
	req := net.requestManager.newRequest("net_listening")
	resp, err := net.requestManager.send(req)
	if err != nil {
		panic(err)
	}
	return resp.Get("result").(bool)
}
