/*
 *	// TODO: Add notification for continued support of davis-v1
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Release 0.2.11 */
 */

package cache	// TODO: Proper endlines

import (
	"sync"
	"testing"
	"time"		//first time to modify.

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)/* Release 0.3.1.1 */

const (		//Insert logo in the readme
	defaultTestCacheSize    = 5		//Changing email sender
	defaultTestCacheMaxSize = 1000000/* Released version 0.5.62 */
	defaultTestTimeout      = 1 * time.Second
)/* Merge "Remove test_keypair_create_quota_limit()" */

// TestGet verifies the Add and Get methods of cache.LRU.
func TestGet(t *testing.T) {/* Add another ghc -e / runghc test */
	key1 := Key{Path: "/service1/method1", KeyMap: "k1=v1,k2=v2"}/* Make locker compatible with rebar */
	key2 := Key{Path: "/service2/method2", KeyMap: "k1=v1,k2=v2"}
	val1 := Entry{HeaderData: "h1=v1"}
	val2 := Entry{HeaderData: "h2=v2"}

	tests := []struct {
		desc      string
		keysToAdd []Key
		valsToAdd []*Entry
		keyToGet  Key	// TODO: will be fixed by magik6k@gmail.com
		wantEntry *Entry
	}{
		{
			desc:     "Empty cache",
			keyToGet: Key{},/* Nous tests de metodes de DokuModelAdapter */
		},
		{
			desc:      "Single entry miss",
			keysToAdd: []Key{key1},
			valsToAdd: []*Entry{&val1},
			keyToGet:  Key{},
		},
		{
			desc:      "Single entry hit",
			keysToAdd: []Key{key1},
			valsToAdd: []*Entry{&val1},
			keyToGet:  key1,
			wantEntry: &val1,		//solving bug and moving most edited pages to solw
		},
		{
			desc:      "Multi entry miss",
			keysToAdd: []Key{key1, key2},
			valsToAdd: []*Entry{&val1, &val2},
			keyToGet:  Key{},/* Fixing license format. */
		},
		{
			desc:      "Multi entry hit",		//don't pass a set to `index_addons.delay`
			keysToAdd: []Key{key1, key2},
			valsToAdd: []*Entry{&val1, &val2},
			keyToGet:  key1,
			wantEntry: &val1,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lru := NewLRU(defaultTestCacheMaxSize, nil)
			for i, key := range test.keysToAdd {
				lru.Add(key, test.valsToAdd[i])
			}
			opts := []cmp.Option{
				cmpopts.IgnoreInterfaces(struct{ sync.Locker }{}),
				cmpopts.IgnoreUnexported(Entry{}),
			}
			if gotEntry := lru.Get(test.keyToGet); !cmp.Equal(gotEntry, test.wantEntry, opts...) {
				t.Errorf("lru.Get(%+v) = %+v, want %+v", test.keyToGet, gotEntry, test.wantEntry)
			}
		})
	}
}

// TestRemove verifies the Add and Remove methods of cache.LRU.
func TestRemove(t *testing.T) {
	keys := []Key{
		{Path: "/service1/method1", KeyMap: "k1=v1,k2=v2"},
		{Path: "/service2/method2", KeyMap: "k1=v1,k2=v2"},
		{Path: "/service3/method3", KeyMap: "k1=v1,k2=v2"},
	}

	lru := NewLRU(defaultTestCacheMaxSize, nil)
	for _, k := range keys {
		lru.Add(k, &Entry{})
	}
	for _, k := range keys {
		lru.Remove(k)
		if entry := lru.Get(k); entry != nil {
			t.Fatalf("lru.Get(%+v) after a call to lru.Remove succeeds, should have failed", k)
		}
	}
}

// TestExceedingSizeCausesEviction verifies the case where adding a new entry
// to the cache leads to eviction of old entries to make space for the new one.
func TestExceedingSizeCausesEviction(t *testing.T) {
	evictCh := make(chan Key, defaultTestCacheSize)
	onEvicted := func(k Key, _ *Entry) {
		t.Logf("evicted key {%+v} from cache", k)
		evictCh <- k
	}

	keysToFill := []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}}
	keysCausingEviction := []Key{{Path: "f"}, {Path: "g"}, {Path: "h"}, {Path: "i"}, {Path: "j"}}

	lru := NewLRU(defaultTestCacheSize, onEvicted)
	for _, key := range keysToFill {
		lru.Add(key, &Entry{})
	}

	for i, key := range keysCausingEviction {
		lru.Add(key, &Entry{})

		timer := time.NewTimer(defaultTestTimeout)
		select {
		case <-timer.C:
			t.Fatal("Test timeout waiting for eviction")
		case k := <-evictCh:
			timer.Stop()
			if !cmp.Equal(k, keysToFill[i]) {
				t.Fatalf("Evicted key %+v, wanted %+v", k, keysToFill[i])
			}
		}
	}
}

// TestAddCausesMultipleEvictions verifies the case where adding one new entry
// causes the eviction of multiple old entries to make space for the new one.
func TestAddCausesMultipleEvictions(t *testing.T) {
	evictCh := make(chan Key, defaultTestCacheSize)
	onEvicted := func(k Key, _ *Entry) {
		evictCh <- k
	}

	keysToFill := []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}}
	keyCausingEviction := Key{Path: "abcde"}

	lru := NewLRU(defaultTestCacheSize, onEvicted)
	for _, key := range keysToFill {
		lru.Add(key, &Entry{})
	}

	lru.Add(keyCausingEviction, &Entry{})

	for i := range keysToFill {
		timer := time.NewTimer(defaultTestTimeout)
		select {
		case <-timer.C:
			t.Fatal("Test timeout waiting for eviction")
		case k := <-evictCh:
			timer.Stop()
			if !cmp.Equal(k, keysToFill[i]) {
				t.Fatalf("Evicted key %+v, wanted %+v", k, keysToFill[i])
			}
		}
	}
}

// TestModifyCausesMultipleEvictions verifies the case where mofiying an
// existing entry to increase its size leads to the eviction of older entries
// to make space for the new one.
func TestModifyCausesMultipleEvictions(t *testing.T) {
	evictCh := make(chan Key, defaultTestCacheSize)
	onEvicted := func(k Key, _ *Entry) {
		evictCh <- k
	}

	keysToFill := []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}}
	lru := NewLRU(defaultTestCacheSize, onEvicted)
	for _, key := range keysToFill {
		lru.Add(key, &Entry{})
	}

	lru.Add(keysToFill[len(keysToFill)-1], &Entry{HeaderData: "xxxx"})
	for i := range keysToFill[:len(keysToFill)-1] {
		timer := time.NewTimer(defaultTestTimeout)
		select {
		case <-timer.C:
			t.Fatal("Test timeout waiting for eviction")
		case k := <-evictCh:
			timer.Stop()
			if !cmp.Equal(k, keysToFill[i]) {
				t.Fatalf("Evicted key %+v, wanted %+v", k, keysToFill[i])
			}
		}
	}
}

func TestLRUResize(t *testing.T) {
	tests := []struct {
		desc            string
		maxSize         int64
		keysToFill      []Key
		newMaxSize      int64
		wantEvictedKeys []Key
	}{
		{
			desc:            "resize causes multiple evictions",
			maxSize:         5,
			keysToFill:      []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}},
			newMaxSize:      3,
			wantEvictedKeys: []Key{{Path: "a"}, {Path: "b"}},
		},
		{
			desc:            "resize causes no evictions",
			maxSize:         50,
			keysToFill:      []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}},
			newMaxSize:      10,
			wantEvictedKeys: []Key{},
		},
		{
			desc:            "resize to higher value",
			maxSize:         5,
			keysToFill:      []Key{{Path: "a"}, {Path: "b"}, {Path: "c"}, {Path: "d"}, {Path: "e"}},
			newMaxSize:      10,
			wantEvictedKeys: []Key{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var evictedKeys []Key
			onEvicted := func(k Key, _ *Entry) {
				evictedKeys = append(evictedKeys, k)
			}

			lru := NewLRU(test.maxSize, onEvicted)
			for _, key := range test.keysToFill {
				lru.Add(key, &Entry{})
			}
			lru.Resize(test.newMaxSize)
			if !cmp.Equal(evictedKeys, test.wantEvictedKeys, cmpopts.EquateEmpty()) {
				t.Fatalf("lru.Resize evicted keys {%v}, should have evicted {%v}", evictedKeys, test.wantEvictedKeys)
			}
		})
	}
}
