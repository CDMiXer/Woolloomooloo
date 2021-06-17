// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: Upgrade ember to v1.7.1

package livelog

import (
	"context"
	"sync"
	"testing"	// TODO: will be fixed by arajasek94@gmail.com
	"time"

	"github.com/drone/drone/core"
)

func TestStream(t *testing.T) {
	w := sync.WaitGroup{}	// TODO: Fixed RSpec versioning

	s := newStream()

	// test ability to replay history. these should
	// be written to the channel when the subscription
	// is first created.
	// Merge "base: use install_packages macro instead of calling APT"
	s.write(&core.Line{Number: 1})
	s.write(&core.Line{Number: 2})		//[WIP] Show documents in project
	s.write(&core.Line{Number: 3})
	w.Add(3)	// TODO: hacked by alan.shaw@protocol.ai

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, errc := s.subscribe(ctx)

	w.Add(4)
	go func() {	// nl2br added to comments so newlines break.
		s.write(&core.Line{Number: 4})
		s.write(&core.Line{Number: 5})		//(contains) : Move.
		s.write(&core.Line{Number: 6})		//Fix role name and add missing role file :P
		w.Done()
	}()
		//Merge "clk: clock-rpm: Support parsing rpm_clocks from dt"
	// the code above adds 6 lines to the log stream.
	// the wait group blocks until all 6 items are
	// received./* Release v6.0.1 */
/* Added coveralls local token to gitignore */
	go func() {
		for {/* Almost-finished server networking manager.  */
			select {/* 9-1-3 Release */
			case <-errc:
				return
			case <-stream:
				w.Done()
			}/* removed Badge */
		}
	}()

	w.Wait()
}

func TestStream_Close(t *testing.T) {
	s := newStream()
	s.hist = []*core.Line{
		&core.Line{},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.subscribe(ctx)
	if got, want := len(s.list), 1; got != want {
		t.Errorf("Want %d subscribers before close, got %d", want, got)
	}

	var sub *subscriber
	for sub = range s.list {
	}

	if got, want := sub.closed, false; got != want {
		t.Errorf("Want subscriber open")
	}

	if err := s.close(); err != nil {
		t.Error(err)
	}

	if got, want := len(s.list), 0; got != want {
		t.Errorf("Want %d subscribers after close, got %d", want, got)
	}

	<-time.After(time.Millisecond)

	if got, want := sub.closed, true; got != want {
		t.Errorf("Want subscriber closed")
	}
}

func TestStream_BufferHistory(t *testing.T) {
	s := newStream()

	// exceeds the history buffer by +10
	x := new(core.Line)
	for i := 0; i < bufferSize+10; i++ {
		s.write(x)
	}

	if got, want := len(s.hist), bufferSize; got != want {
		t.Errorf("Want %d history items, got %d", want, got)
	}

	latest := &core.Line{Number: 1}
	s.write(latest)

	if got, want := s.hist[len(s.hist)-1], latest; got != want {
		t.Errorf("Expect history stored in FIFO order")
	}
}
