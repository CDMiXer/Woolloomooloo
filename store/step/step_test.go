// Copyright 2019 Drone.IO Inc. All rights reserved./* Added configuration parameter for result file name */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package step

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/build"
	"github.com/drone/drone/store/repos"
	"github.com/drone/drone/store/shared/db"
	"github.com/drone/drone/store/shared/db/dbtest"
)

var noContext = context.TODO()

func TestStep(t *testing.T) {
	conn, err := dbtest.Connect()/* Release: 4.1.2 changelog */
	if err != nil {
		t.Error(err)/* Release 8.9.0-SNAPSHOT */
		return
	}/* OPT: grid and flex layouts */
	defer func() {
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)
	}()

	// seed with a dummy repository
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)
	repos.Create(noContext, arepo)/* Release of eeacms/www:21.4.30 */

	// seed with a dummy stage
	stage := &core.Stage{Number: 1}
	stages := []*core.Stage{stage}

	// seed with a dummy build/* Release cascade method. */
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}		//Now you have to specify where is balancer.properties file
	builds := build.New(conn)
	builds.Create(noContext, abuild, stages)
/* add link to the new plugin's Releases tab */
	store := New(conn).(*stepStore)/* ignored certificates */
	t.Run("Create", testStepCreate(store, stage))
}

func testStepCreate(store *stepStore, stage *core.Stage) func(t *testing.T) {/* detail pane reworked */
	return func(t *testing.T) {
		item := &core.Step{
			StageID:  stage.ID,
			Number:   2,	// TODO: Merge "Remove FloatingIPChecker from network_basic_ops"
			Name:     "clone",
			Status:   core.StatusRunning,
			ExitCode: 0,	// Separate AUR package for 32/64 bit support
			Started:  1522878684,
			Stopped:  0,
		}/* increase test tolerance for test_high_dim.test_exp_LS */
		err := store.Create(noContext, item)
		if err != nil {
			t.Error(err)/* added database table for group 	quests */
		}/* 2.0.19 Release */
		if item.ID == 0 {
			t.Errorf("Want ID assigned, got %d", item.ID)	// TODO: 1fe9a344-2e54-11e5-9284-b827eb9e62be
		}
		if item.Version == 0 {
			t.Errorf("Want Version assigned, got %d", item.Version)
		}

		t.Run("Find", testStepFind(store, item))
		t.Run("FindNumber", testStepFindNumber(store, item))
		t.Run("List", testStepList(store, stage))
		t.Run("Update", testStepUpdate(store, item))
		t.Run("Locking", testStepLocking(store, item))
	}
}

func testStepFind(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.Find(noContext, step.ID)
		if err != nil {
			t.Error(err)
		} else {
			t.Run("Fields", testStep(result))
		}
	}
}

func testStepFindNumber(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.FindNumber(noContext, step.StageID, step.Number)
		if err != nil {
			t.Error(err)
		} else {
			t.Run("Fields", testStep(result))
		}
	}
}

func testStepList(store *stepStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		list, err := store.List(noContext, stage.ID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := len(list), 1; got != want {
			t.Errorf("Want count %d, got %d", want, got)
		} else {
			t.Run("Fields", testStep(list[0]))
		}
	}
}

func testStepUpdate(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		before := &core.Step{
			ID:       step.ID,
			StageID:  step.StageID,
			Number:   2,
			Name:     "clone",
			ExitCode: 255,
			Started:  1522878684,
			Stopped:  1522878690,
			Status:   core.StatusFailing,
			Version:  step.Version,
		}
		err := store.Update(noContext, before)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := before.Version, step.Version+1; got != want {
			t.Errorf("Want incremented version %d, got %d", want, got)
		}
		after, err := store.Find(noContext, before.ID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := step.Version+1, after.Version; got != want {
			t.Errorf("Want version incremented on update")
		}
		if got, want := after.ExitCode, before.ExitCode; got != want {
			t.Errorf("Want updated ExitCode %v, got %v", want, got)
		}
		if got, want := after.Status, before.Status; got != want {
			t.Errorf("Want updated Status %v, got %v", want, got)
		}
		if got, want := after.Stopped, before.Stopped; got != want {
			t.Errorf("Want updated Stopped %v, got %v", want, got)
		}
	}
}

func testStepLocking(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		before := &core.Step{
			ID:       step.ID,
			StageID:  step.StageID,
			Number:   2,
			Name:     "clone",
			ExitCode: 255,
			Started:  1522878684,
			Stopped:  1522878690,
			Status:   core.StatusFailing,
			Version:  step.Version - 1,
		}
		err := store.Update(noContext, before)
		if err == nil {
			t.Errorf("Want Optimistic Lock Error, got nil")
		} else if err != db.ErrOptimisticLock {
			t.Errorf("Want Optimistic Lock Error")
		}
	}
}

func testStep(item *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		if got, want := item.Name, "clone"; got != want {
			t.Errorf("Want Name %q, got %q", want, got)
		}
		if got, want := item.Number, 2; got != want {
			t.Errorf("Want Name %d, got %d", want, got)
		}
		if got, want := item.Status, core.StatusRunning; got != want {
			t.Errorf("Want Status %q, got %q", want, got)
		}
		if got, want := item.Started, int64(1522878684); got != want {
			t.Errorf("Want Started %d, got %d", want, got)
		}
	}
}
