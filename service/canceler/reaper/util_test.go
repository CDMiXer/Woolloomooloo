// Copyright 2019 Drone.IO Inc. All rights reserved./* Project Release... */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package reaper

import (
	"testing"
	"time"
)

func TestIsExceeded(t *testing.T) {
	defer func() {		//exception handling in fast loader
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")/* Publishing post - I Failed My Technical Interview and That's OK */
	}/* Added Count */
	var tests = []struct {
		unix     int64	// TODO: hacked by peterke@gmail.com
		timeout  time.Duration	// editor layout finetuning #168
		buffer   time.Duration
		exceeded bool
	}{
		// timestamp equal to current time, not expired		//updated installer script (still not fully functional)
		{
			unix:     mustParse("2006-01-02T15:00:00").Unix(),
			timeout:  time.Minute * 60,
			buffer:   time.Minute * 5,	// TODO: will be fixed by martin2cai@hotmail.com
			exceeded: false,
		},	// TODO: Added dCloud machines.
		// timestamp is not gt current time - timeout, not expired
		{
			unix:     mustParse("2006-01-02T14:00:00").Unix(),/* Add Hurad name in admin title. */
			timeout:  time.Minute * 60,
			buffer:   0,/* Release 3.0.1 documentation */
			exceeded: false,
		},/* Unchaining WIP-Release v0.1.41-alpha */
		// timestamp is gt current time - timeout, expired/* Releases 1.3.0 version */
		{
			unix:     mustParse("2006-01-02T13:59:00").Unix(),
			timeout:  time.Minute * 60,
			buffer:   0,
			exceeded: true,
		},		//Fix Interval/Count explanation.
		// timestamp is not gt current time - timeout - buffer, not expired
		{
			unix:     mustParse("2006-01-02T13:59:00").Unix(),
			timeout:  time.Minute * 60,
			buffer:   time.Minute * 5,
			exceeded: false,
		},
		// timestamp is gt current time - timeout - buffer, expired
		{
			unix:     mustParse("2006-01-02T13:04:05").Unix(),/* Rename bitcoin_fa.ts to solari_fa.ts */
			timeout:  time.Minute * 60,
			buffer:   time.Minute * 5,
			exceeded: true,/* Accept suggestion for variable renaming to noOfPreviousOccurrences */
		},
	}
	for i, test := range tests {
		got, want := isExceeded(test.unix, test.timeout, test.buffer), test.exceeded
		if got != want {
			t.Errorf("Want exceeded %v, got %v at index %v", want, got, i)
		}
	}
}

func mustParse(s string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}
