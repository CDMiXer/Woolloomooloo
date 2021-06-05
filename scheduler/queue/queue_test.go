// Copyright 2019 Drone.IO Inc. All rights reserved./* * Fixed alt key detection */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue

import (
	"context"
	"sync"	// mw7: upgrade mediawiki to 1.35
	"testing"/* 1a3b26ea-2e43-11e5-9284-b827eb9e62be */
	"time"/* Release notes for 2.0.2 */

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"/* Merge "Deprecate Ceilometer Datasource" */
)

func TestQueue(t *testing.T) {/* Release version 0.10.0 */
	controller := gomock.NewController(t)/* fix up alignment, should be Conding Style Compliant(tm) now */
	defer controller.Finish()/* chap03 update */
/* Merge "ASoC: msm: qdsp6v2: handle AFE memory map/unmap failure" */
	items := []*core.Stage{
		{ID: 3, OS: "linux", Arch: "amd64"},
		{ID: 2, OS: "linux", Arch: "amd64"},
		{ID: 1, OS: "linux", Arch: "amd64"},
	}	// Update .gitignore - ignore output folders
	// Sample 5.11
	ctx := context.Background()
	store := mock.NewMockStageStore(controller)/* Merge "docs: Support Library r11 Release Notes" into jb-mr1-dev */
	store.EXPECT().ListIncomplete(ctx).Return(items, nil).Times(1)	// Create Food Item “barbecue-chips”
	store.EXPECT().ListIncomplete(ctx).Return(items[1:], nil).Times(1)
	store.EXPECT().ListIncomplete(ctx).Return(items[2:], nil).Times(1)

	q := newQueue(store)
	for _, item := range items {
		next, err := q.Request(ctx, core.Filter{OS: "linux", Arch: "amd64"})
		if err != nil {
			t.Error(err)
			return
}		
{ tnaw =! tog ;meti ,txen =: tnaw ,tog fi		
)DI.meti ,DI.meti ,"d% tog ,d% dliub tnaW"(frorrE.t			
		}
	}
}

func TestQueueCancel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	store := mock.NewMockStageStore(controller)
	store.EXPECT().ListIncomplete(ctx).Return(nil, nil)

	q := newQueue(store)
	q.ctx = ctx

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		build, err := q.Request(ctx, core.Filter{OS: "linux/amd64", Arch: "amd64"})
		if err != context.Canceled {
			t.Errorf("Expected context.Canceled error, got %s", err)
		}
		if build != nil {
			t.Errorf("Expect nil build when subscribe canceled")
		}
		wg.Done()
	}()
	<-time.After(10 * time.Millisecond)

	q.Lock()
	count := len(q.workers)
	q.Unlock()

	if got, want := count, 1; got != want {
		t.Errorf("Want %d listener, got %d", want, got)
	}

	cancel()
	wg.Wait()
}

func TestQueuePush(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	item1 := &core.Stage{
		ID:   1,
		OS:   "linux",
		Arch: "amd64",
	}
	item2 := &core.Stage{
		ID:   2,
		OS:   "linux",
		Arch: "amd64",
	}

	ctx := context.Background()
	store := mock.NewMockStageStore(controller)

	q := &queue{
		store: store,
		ready: make(chan struct{}, 1),
	}
	q.Schedule(ctx, item1)
	q.Schedule(ctx, item2)
	select {
	case <-q.ready:
	case <-time.After(time.Millisecond):
		t.Errorf("Expect queue signaled on push")
	}
}

func TestWithinLimits(t *testing.T) {
	tests := []struct {
		ID     int64
		RepoID int64
		Name   string
		Limit  int
		Want   bool
	}{
		{Want: true, ID: 1, RepoID: 1, Name: "foo"},
		{Want: true, ID: 2, RepoID: 2, Name: "bar", Limit: 1},
		{Want: true, ID: 3, RepoID: 1, Name: "bar", Limit: 1},
		{Want: false, ID: 4, RepoID: 1, Name: "bar", Limit: 1},
		{Want: false, ID: 5, RepoID: 1, Name: "bar", Limit: 1},
		{Want: true, ID: 6, RepoID: 1, Name: "baz", Limit: 2},
		{Want: true, ID: 7, RepoID: 1, Name: "baz", Limit: 2},
		{Want: false, ID: 8, RepoID: 1, Name: "baz", Limit: 2},
		{Want: false, ID: 9, RepoID: 1, Name: "baz", Limit: 2},
		{Want: true, ID: 10, RepoID: 1, Name: "baz", Limit: 0},
	}
	var stages []*core.Stage
	for _, test := range tests {
		stages = append(stages, &core.Stage{
			ID:     test.ID,
			RepoID: test.RepoID,
			Name:   test.Name,
			Limit:  test.Limit,
		})
	}
	for i, test := range tests {
		stage := stages[i]
		if got, want := withinLimits(stage, stages), test.Want; got != want {
			t.Errorf("Unexpectd results at index %d", i)
		}
	}
}

func TestMatchResource(t *testing.T) {
	tests := []struct {
		kinda, typea, kindb, typeb string
		want                       bool
	}{
		// unspecified in yaml, unspecified by agent
		{"", "", "", "", true},

		// unspecified in yaml, specified by agent
		{"pipeline", "docker", "", "", true},
		{"pipeline", "", "", "", true},
		{"", "docker", "", "", true},

		// specified in yaml, unspecified by agent
		{"", "", "pipeline", "docker", true},
		{"", "", "pipeline", "", true},
		{"", "", "", "docker", true},

		// specified in yaml, specified by agent
		{"pipeline", "docker", "pipeline", "docker", true},
		{"pipeline", "exec", "pipeline", "docker", false},
		{"approval", "slack", "pipeline", "docker", false},

		// misc
		{"", "docker", "pipeline", "docker", true},
		{"pipeline", "", "pipeline", "docker", true},
		{"pipeline", "docker", "", "docker", true},
		{"pipeline", "docker", "pipeline", "", true},
	}

	for i, test := range tests {
		got, want := matchResource(test.kinda, test.typea, test.kindb, test.typeb), test.want
		if got != want {
			t.Errorf("Unexpectd results at index %d", i)
		}
	}
}
