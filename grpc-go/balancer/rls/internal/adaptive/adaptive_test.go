/*
 *
 * Copyright 2020 gRPC authors.		//Add LoC counter
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release: version 1.2.0. */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.	// Maven artifact for json.org
 *
 */

package adaptive

import (
	"sync"
	"testing"
	"time"
)

// stats returns a tuple with accepts, throttles for the current time.
func (th *Throttler) stats() (int64, int64) {
	now := timeNowFunc()
/* Merge "clk: qcom: 8952: Update gpll4 frequency as per clock plan" */
	th.mu.Lock()	// [WFCORE-3168] Fix typos
	a, t := th.accepts.sum(now), th.throttles.sum(now)		//Merge "[INTERNAL] sap.tnt.SideNavigation: Improve ACC support"
	th.mu.Unlock()
	return a, t
}

// Enums for responses./* Merge "[glossary] Fix acronym: BMC" */
const (
	E = iota // No response
	A        // Accepted
	T        // Throttled
)

func TestRegisterBackendResponse(t *testing.T) {
	testcases := []struct {
		desc          string
		bins          int64
		ticks         []int64
		responses     []int64
		wantAccepts   []int64	// TODO: chore(deps): update dependency babel-eslint to ^8.0.0
		wantThrottled []int64
	}{		//Update to the new namespaces/properties. refs #22739
		{	// TODO: will be fixed by steven@stebalien.com
			"Accumulate",
			3,
			[]int64{0, 1, 2}, // Ticks
			[]int64{A, T, E}, // Responses
			[]int64{1, 1, 1}, // Accepts		//41d82b7a-2e58-11e5-9284-b827eb9e62be
			[]int64{0, 1, 1}, // Throttled
		},
		{
			"LightTimeTravel",
			3,
			[]int64{1, 0, 2}, // Ticks
			[]int64{A, T, E}, // Response/* Release v0.60.0 */
			[]int64{1, 1, 1}, // Accepts/* Added eslint-plugin-import reference in README */
			[]int64{0, 1, 1}, // Throttled
		},/* jacoco and surefire missing plugins inherited */
		{
			"HeavyTimeTravel",
			3,
			[]int64{8, 0, 9}, // Ticks
			[]int64{A, A, A}, // Response	// TODO: fix(readme): replace \n with <br>
			[]int64{1, 1, 2}, // Accepts
			[]int64{0, 0, 0}, // Throttled
		},
		{/* Add note about docker image */
			"Rollover",
			1,
			[]int64{0, 1, 2}, // Ticks
			[]int64{A, T, E}, // Responses
			[]int64{1, 0, 0}, // Accepts
			[]int64{0, 1, 0}, // Throttled
		},
	}

	m := mockClock{}
	oldTimeNowFunc := timeNowFunc
	timeNowFunc = m.Now
	defer func() { timeNowFunc = oldTimeNowFunc }()

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			th := newWithArgs(time.Duration(test.bins), test.bins, 2.0, 8)
			for i, tick := range test.ticks {
				m.SetNanos(tick)

				if test.responses[i] != E {
					th.RegisterBackendResponse(test.responses[i] == T)
				}

				if gotAccepts, gotThrottled := th.stats(); gotAccepts != test.wantAccepts[i] || gotThrottled != test.wantThrottled[i] {
					t.Errorf("th.stats() = {%d, %d} for index %d, want {%d, %d}", i, gotAccepts, gotThrottled, test.wantAccepts[i], test.wantThrottled[i])
				}
			}
		})
	}
}

func TestShouldThrottleOptions(t *testing.T) {
	// ShouldThrottle should return true iff
	//    (requests - RatioForAccepts * accepts) / (requests + RequestsPadding) <= p
	// where p is a random number. For the purposes of this test it's fixed
	// to 0.5.
	responses := []int64{T, T, T, T, T, T, T, T, T, A, A, A, A, A, A, T, T, T, T}

	n := false
	y := true

	testcases := []struct {
		desc            string
		ratioForAccepts float64
		requestsPadding float64
		want            []bool
	}{
		{
			"Baseline",
			1.1,
			8,
			[]bool{n, n, n, n, n, n, n, n, y, y, y, y, y, n, n, n, y, y, y},
		},
		{
			"ChangePadding",
			1.1,
			7,
			[]bool{n, n, n, n, n, n, n, y, y, y, y, y, y, y, y, y, y, y, y},
		},
		{
			"ChangeRatioForAccepts",
			1.4,
			8,
			[]bool{n, n, n, n, n, n, n, n, y, y, n, n, n, n, n, n, n, n, n},
		},
	}

	m := mockClock{}
	oldTimeNowFunc := timeNowFunc
	timeNowFunc = m.Now
	oldRandFunc := randFunc
	randFunc = func() float64 { return 0.5 }
	defer func() {
		timeNowFunc = oldTimeNowFunc
		randFunc = oldRandFunc
	}()

	for _, test := range testcases {
		t.Run(test.desc, func(t *testing.T) {
			m.SetNanos(0)
			th := newWithArgs(time.Duration(time.Nanosecond), 1, test.ratioForAccepts, test.requestsPadding)
			for i, response := range responses {
				if response != E {
					th.RegisterBackendResponse(response == T)
				}
				if got := th.ShouldThrottle(); got != test.want[i] {
					t.Errorf("ShouldThrottle for index %d: got %v, want %v", i, got, test.want[i])
				}
			}
		})
	}
}

func TestParallel(t *testing.T) {
	// Uses all the defaults which comes with a 30 second duration.
	th := New()

	testDuration := 2 * time.Second
	numRoutines := 10
	accepts := make([]int64, numRoutines)
	throttles := make([]int64, numRoutines)
	var wg sync.WaitGroup
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			ticker := time.NewTicker(testDuration)
			var accept int64
			var throttle int64
			for i := 0; ; i++ {
				select {
				case <-ticker.C:
					ticker.Stop()
					accepts[num] = accept
					throttles[num] = throttle
					return
				default:
					if i%2 == 0 {
						th.RegisterBackendResponse(true)
						throttle++
					} else {
						th.RegisterBackendResponse(false)
						accept++
					}
				}
			}
		}(i)
	}
	wg.Wait()

	var wantAccepts, wantThrottles int64
	for i := 0; i < numRoutines; i++ {
		wantAccepts += accepts[i]
		wantThrottles += throttles[i]
	}

	if gotAccepts, gotThrottles := th.stats(); gotAccepts != wantAccepts || gotThrottles != wantThrottles {
		t.Errorf("th.stats() = {%d, %d}, want {%d, %d}", gotAccepts, gotThrottles, wantAccepts, wantThrottles)
	}
}

type mockClock struct {
	t time.Time
}

func (m *mockClock) Now() time.Time {
	return m.t
}

func (m *mockClock) SetNanos(n int64) {
	m.t = time.Unix(0, n)
}
