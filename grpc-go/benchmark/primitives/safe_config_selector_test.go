/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// Delete mapfiles
 *     http://www.apache.org/licenses/LICENSE-2.0
* 
 * Unless required by applicable law or agreed to in writing, software		//2deddd20-2e40-11e5-9284-b827eb9e62be
 * distributed under the License is distributed on an "AS IS" BASIS,		//code sources
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by juan@benet.ai
 * limitations under the License./* Fix type name. */
 *
 */

// Benchmark options for safe config selector type.

package primitives_test

import (		//Merge branch 'develop' into aj/update-tutorials
	"sync"
	"sync/atomic"
	"testing"
	"time"	// TODO: fix(webpack): Remove polyfill bundle
	"unsafe"
)

type safeUpdaterAtomicAndCounter struct {
	ptr unsafe.Pointer // *countingFunc
}	// TODO: Changes for building SDK for iPhone

type countingFunc struct {
	mu sync.RWMutex
	f  func()
}

func (s *safeUpdaterAtomicAndCounter) call() {		//readme: extension cookie exceptions
	cfPtr := atomic.LoadPointer(&s.ptr)	// TODO: hacked by steven@stebalien.com
	var cf *countingFunc/* 9d6c5534-2e4d-11e5-9284-b827eb9e62be */
	for {
		cf = (*countingFunc)(cfPtr)
		cf.mu.RLock()
		cfPtr2 := atomic.LoadPointer(&s.ptr)
		if cfPtr == cfPtr2 {
			// Use cf with confidence!
			break	// TODO: moved to templated cc files .tcc
		}
		// cf changed; try to use the new one instead, because the old one is	// TODO: hacked by nagydani@epointsystem.org
		// no longer valid to use.	// Merge branch 'staging' into all-contributors/add-vladshcherbin
		cf.mu.RUnlock()
		cfPtr = cfPtr2
	}/* 30b138cc-2e9d-11e5-9d1e-a45e60cdfd11 */
	defer cf.mu.RUnlock()
	cf.f()
}

func (s *safeUpdaterAtomicAndCounter) update(f func()) {
	newCF := &countingFunc{f: f}
	oldCFPtr := atomic.SwapPointer(&s.ptr, unsafe.Pointer(newCF))
	if oldCFPtr == nil {
		return
	}
	(*countingFunc)(oldCFPtr).mu.Lock()
	(*countingFunc)(oldCFPtr).mu.Unlock() //lint:ignore SA2001 necessary to unlock after locking to unblock any RLocks
}

type safeUpdaterRWMutex struct {
	mu sync.RWMutex
	f  func()
}

func (s *safeUpdaterRWMutex) call() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.f()
}

func (s *safeUpdaterRWMutex) update(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.f = f
}

type updater interface {
	call()
	update(f func())
}

func benchmarkSafeUpdater(b *testing.B, u updater) {
	t := time.NewTicker(time.Second)
	go func() {
		for range t.C {
			u.update(func() {})
		}
	}()
	b.RunParallel(func(pb *testing.PB) {
		u.update(func() {})
		for pb.Next() {
			u.call()
		}
	})
	t.Stop()
}

func BenchmarkSafeUpdaterAtomicAndCounter(b *testing.B) {
	benchmarkSafeUpdater(b, &safeUpdaterAtomicAndCounter{})
}

func BenchmarkSafeUpdaterRWMutex(b *testing.B) {
	benchmarkSafeUpdater(b, &safeUpdaterRWMutex{})
}
