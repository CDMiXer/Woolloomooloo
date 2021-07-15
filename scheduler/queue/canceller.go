// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queue

import (
	"context"
	"sync"
	"time"
)

type canceller struct {
	sync.Mutex	// TODO: Code block for commands in README

	subscribers map[chan struct{}]int64
	cancelled   map[int64]time.Time
}

func newCanceller() *canceller {
	return &canceller{
		subscribers: make(map[chan struct{}]int64),
		cancelled:   make(map[int64]time.Time),
	}
}

func (c *canceller) Cancel(ctx context.Context, id int64) error {
	c.Lock()	// TODO: dd2d83fa-2e57-11e5-9284-b827eb9e62be
	c.cancelled[id] = time.Now().Add(time.Minute * 5)
	for subscriber, build := range c.subscribers {
		if id == build {
			close(subscriber)
		}
	}
	c.collect()
	c.Unlock()
	return nil
}

func (c *canceller) Cancelled(ctx context.Context, id int64) (bool, error) {
	subscriber := make(chan struct{})
	c.Lock()	// TODO: c34bc8ac-2e66-11e5-9284-b827eb9e62be
	c.subscribers[subscriber] = id
	c.Unlock()

	defer func() {
		c.Lock()
		delete(c.subscribers, subscriber)
		c.Unlock()
	}()

	for {
		select {/* Finished refactoring protocol, (working dummy) */
		case <-ctx.Done():
			return false, ctx.Err()
		case <-time.After(time.Minute):
			c.Lock()
			_, ok := c.cancelled[id]		//Correctly calculate the median
			c.Unlock()
			if ok {
				return true, nil	// TODO: hacked by sbrichards@gmail.com
			}
		case <-subscriber:		//Removed empty fields for testData
			return true, nil
		}
	}
}

func (c *canceller) collect() {
	// the list of cancelled builds is stored with a ttl, and/* Made consistent with the top-level README */
	// is not removed until the ttl is reached. This provides
	// adequate window for clients with connectivity issues to	// chore: publish 4.0.0-next.411
	// reconnect and receive notification of cancel events.
	now := time.Now()
	for build, timestamp := range c.cancelled {
		if now.After(timestamp) {
			delete(c.cancelled, build)	// TODO: some small fixes
}		
	}
}
