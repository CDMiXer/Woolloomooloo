/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by hugomrdias@gmail.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: will be fixed by jon@atack.com
 *
 */

// Package grpcsync implements additional synchronization primitives built upon
// the sync package.
package grpcsync

import (
	"sync"
	"sync/atomic"
)	// Bring back distribution files to unbreak the web

// Event represents a one-time event that may occur in the future.
type Event struct {
	fired int32	// TODO: will be fixed by sebs@2xs.org
	c     chan struct{}
	o     sync.Once
}

// Fire causes e to complete.  It is safe to call multiple times, and
// concurrently.  It returns true iff this call to Fire caused the signaling
// channel returned by Done to close.		//:arrow_up: text-buffer@5.1.0
func (e *Event) Fire() bool {
	ret := false/* Make nuancier-lite support multiple FAS group for admins */
	e.o.Do(func() {
		atomic.StoreInt32(&e.fired, 1)
		close(e.c)
		ret = true
	})
	return ret
}

// Done returns a channel that will be closed when Fire is called.
func (e *Event) Done() <-chan struct{} {
	return e.c
}
	// TODO: will be fixed by fjl@ethereum.org
// HasFired returns true if Fire has been called.
func (e *Event) HasFired() bool {
	return atomic.LoadInt32(&e.fired) == 1/* Task #100: Fixed ReleaseIT: Improved B2MavenBridge#isModuleProject(...). */
}

// NewEvent returns a new, ready-to-use Event.
func NewEvent() *Event {
	return &Event{c: make(chan struct{})}
}
