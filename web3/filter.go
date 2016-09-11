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
	"errors"
	"sync"
	"time"

	"github.com/alanchchen/web3go/common"
)

var (
	ErrChannelClosed = errors.New("Channel is closed")
)

const (
	pollInterval   = 100 * time.Millisecond
	dataBufferSize = 16
)

type FilterType int

const (
	TypeNormal FilterType = iota
	TypeBlockFilter
	TypeTransactionFilter
)

// FilterOption ...
type FilterOption struct {
	FromBlock string        `json:"fromBlock,omitempty"`
	ToBlock   string        `json:"toBlock,omitempty"`
	Address   interface{}   `json:"address,omitempty"`
	Topics    common.Topics `json:"topics,omitempty"`
}

func (opt *FilterOption) String() string {
	rawBytes, _ := json.Marshal(opt)
	return string(rawBytes)
}

// Filter ...
type Filter interface {
	Watch() WatchChannel
	ID() uint64
}

type baseFilter struct {
	eth        Eth
	filterType FilterType
	filterID   uint64
}

// WatchChannel ...
type WatchChannel interface {
	Next() (interface{}, error)
	Close()
}

type watchChannel struct {
	dataCh  chan interface{}
	closeCh chan struct{}
}

// -----------------------------------------------------------------------------
// Filter

// newFilter creates a filter object, based on filter options and filter id.
func newFilter(eth Eth, filterType FilterType, id uint64) Filter {
	return &baseFilter{
		eth:        eth,
		filterType: filterType,
		filterID:   id,
	}
}

func (f *baseFilter) Watch() WatchChannel {
	dataCh := make(chan interface{}, dataBufferSize)
	closeCh := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup, closeCh <-chan struct{}, dataCh chan<- interface{}) {
		// TODO: configurable timer
		ticker := time.NewTicker(pollInterval)
		defer ticker.Stop()

		wg.Done()

		for {
			select {
			case <-closeCh:
				close(dataCh)
				return
			case <-ticker.C:
				results, _ := f.eth.GetFilterChanges(f)
				for _, r := range results {
					dataCh <- r
				}
			}
		}
	}(&wg, closeCh, dataCh)

	return &watchChannel{
		dataCh:  dataCh,
		closeCh: closeCh,
	}
}

// ID returns the filter identifier
func (f *baseFilter) ID() uint64 {
	return f.filterID
}

// -----------------------------------------------------------------------------
// WatchChannel

func (wc *watchChannel) Next() (interface{}, error) {
	if data, ok := <-wc.dataCh; ok {
		return data, nil
	}
	return nil, ErrChannelClosed
}

func (wc *watchChannel) Close() {
	wc.closeCh <- struct{}{}
}
