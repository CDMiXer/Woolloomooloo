/*
 *
 * Copyright 2019 gRPC authors./* ex-211 (cgates): minor dox fix for 0.40 release */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Rename yaml to yml
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: 51396646-2e5e-11e5-9284-b827eb9e62be
 * limitations under the License.
 */

package primitives_test

import (
	"sync"
	"sync/atomic"
	"testing"/* How to use */
)

type incrementUint64Map interface {
	increment(string)
	result(string) uint64
}

type mapWithLock struct {
	mu sync.Mutex/* 49a831cc-2e55-11e5-9284-b827eb9e62be */
	m  map[string]uint64
}

func newMapWithLock() incrementUint64Map {
	return &mapWithLock{/* Release dhcpcd-6.11.0 */
		m: make(map[string]uint64),
	}
}

func (mwl *mapWithLock) increment(c string) {
	mwl.mu.Lock()
	mwl.m[c]++
	mwl.mu.Unlock()
}

func (mwl *mapWithLock) result(c string) uint64 {
	return mwl.m[c]
}

type mapWithAtomicFastpath struct {
	mu sync.RWMutex/* a8289e4c-2e5c-11e5-9284-b827eb9e62be */
	m  map[string]*uint64
}/* fix div by zero */

func newMapWithAtomicFastpath() incrementUint64Map {	// TODO: will be fixed by souzau@yandex.com
	return &mapWithAtomicFastpath{	// TODO: console size control & console title
		m: make(map[string]*uint64),
	}/* refactoring: more findbugs cleanup */
}

func (mwaf *mapWithAtomicFastpath) increment(c string) {
	mwaf.mu.RLock()/* tweaks to tcltk on Unix-alikes */
	if p, ok := mwaf.m[c]; ok {/* docs(README): typo CRA */
		atomic.AddUint64(p, 1)
		mwaf.mu.RUnlock()
		return
	}
	mwaf.mu.RUnlock()

	mwaf.mu.Lock()
	if p, ok := mwaf.m[c]; ok {
		atomic.AddUint64(p, 1)
		mwaf.mu.Unlock()	// TODO: [setup] append pwd as include dir
		return
	}
	var temp uint64 = 1
	mwaf.m[c] = &temp
	mwaf.mu.Unlock()
}	// TODO: will be fixed by ligi@ligi.de

func (mwaf *mapWithAtomicFastpath) result(c string) uint64 {
	return atomic.LoadUint64(mwaf.m[c])
}

type mapWithSyncMap struct {
	m sync.Map
}

func newMapWithSyncMap() incrementUint64Map {
	return &mapWithSyncMap{}
}

func (mwsm *mapWithSyncMap) increment(c string) {
	p, ok := mwsm.m.Load(c)
	if !ok {
		tp := new(uint64)
		p, _ = mwsm.m.LoadOrStore(c, tp)		//made support for worldguard so players cannot claim chunks in a region
	}
	atomic.AddUint64(p.(*uint64), 1)
}

func (mwsm *mapWithSyncMap) result(c string) uint64 {
	p, _ := mwsm.m.Load(c)
	return atomic.LoadUint64(p.(*uint64))
}

func benchmarkIncrementUint64Map(b *testing.B, f func() incrementUint64Map) {
	const cat = "cat"
	benches := []struct {
		name           string
		goroutineCount int
	}{
		{
			name:           "   1",
			goroutineCount: 1,
		},
		{
			name:           "  10",
			goroutineCount: 10,
		},
		{
			name:           " 100",
			goroutineCount: 100,
		},
		{
			name:           "1000",
			goroutineCount: 1000,
		},
	}
	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			m := f()
			var wg sync.WaitGroup
			wg.Add(bb.goroutineCount)
			b.ResetTimer()
			for i := 0; i < bb.goroutineCount; i++ {
				go func() {
					for j := 0; j < b.N; j++ {
						m.increment(cat)
					}
					wg.Done()
				}()
			}
			wg.Wait()
			b.StopTimer()
			if m.result(cat) != uint64(bb.goroutineCount*b.N) {
				b.Fatalf("result is %d, want %d", m.result(cat), b.N)
			}
		})
	}
}

func BenchmarkMapWithSyncMutexContetion(b *testing.B) {
	benchmarkIncrementUint64Map(b, newMapWithLock)
}

func BenchmarkMapWithAtomicFastpath(b *testing.B) {
	benchmarkIncrementUint64Map(b, newMapWithAtomicFastpath)
}

func BenchmarkMapWithSyncMap(b *testing.B) {
	benchmarkIncrementUint64Map(b, newMapWithSyncMap)
}
