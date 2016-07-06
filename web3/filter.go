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
	"sync"
	"sync/atomic"
	"time"

	"github.com/alanchchen/web3go/common"
)

const (
	pollInterval = 100 * time.Millisecond
)

// FilterOption ...
type FilterOption struct {
	FromBlock string        `json:"fromBlock"`
	ToBlock   string        `json:"toBlock"`
	Address   interface{}   `json:"address"`
	Topics    common.Topics `json:"topics"`
}

// Filter ...
type Filter interface {
	Watch(func([]common.Log))
	StopWatching()
	GetOption() *FilterOption
	ID() uint64
}

// baseFilter ...
type baseFilter struct {
	eth          Eth
	option       *FilterOption
	filterID     uint64
	callbackLock sync.RWMutex
	callbacks    []func([]common.Log)
	updateCh     chan struct{}
	running      int32
}

// NewFilter creates a filter object, based on filter options and filter id.
func newFilter(eth Eth, option *FilterOption, id uint64) Filter {
	return &baseFilter{
		eth:      eth,
		option:   option,
		filterID: id,
		updateCh: make(chan struct{}),
	}
}

func (f *baseFilter) pollLoop(ready chan bool) {
	atomic.StoreInt32(&f.running, 1)
	defer atomic.StoreInt32(&f.running, 0)
	ready <- true
	// TODO: configurable timer
	timer := time.NewTimer(pollInterval)
	defer timer.Stop()
	for {
		select {
		case <-f.updateCh:
			f.callbackLock.Lock()
			f.callbacks = f.callbacks[:0]
			f.callbackLock.Unlock()
			return
		case <-timer.C:
			f.poll()
		}
	}
}

func (f *baseFilter) poll() {
	if results, err := f.eth.GetFilterChanges(f); err != nil {
		var wg sync.WaitGroup
		f.callbackLock.RLock()
		for _, callback := range f.callbacks {
			wg.Add(1)
			cb := callback
			go func(cb func([]common.Log)) {
				defer wg.Done()
				cb(results)
			}(cb)
		}
		f.callbackLock.RUnlock()
		wg.Wait()
	} else {
		// error
	}
}

func (f *baseFilter) Watch(callback func([]common.Log)) {
	f.callbackLock.Lock()
	f.callbacks = append(f.callbacks, callback)
	f.callbackLock.Unlock()

	if atomic.LoadInt32(&f.running) == 0 {
		ready := make(chan bool)
		go f.pollLoop(ready)
		<-ready
		close(ready)
	}
}

func (f *baseFilter) StopWatching() {
	f.updateCh <- struct{}{}
}

// ID returns the filter identifier
func (f *baseFilter) ID() uint64 {
	return f.filterID
}

// GetOption returns the filter option to the filter
func (f *baseFilter) GetOption() *FilterOption {
	return f.option
}
