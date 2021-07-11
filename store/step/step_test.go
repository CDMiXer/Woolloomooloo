// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* 6f38166e-2e6e-11e5-9284-b827eb9e62be */
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
	conn, err := dbtest.Connect()
	if err != nil {
		t.Error(err)
		return
	}/* Rename ReleaseNotes to ReleaseNotes.md */
	defer func() {
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)		//Logic error if no LIBZ or LIBBZ2 exists 
	}()

	// seed with a dummy repository	// TODO: Update Saxon XLST processor
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)/* Added the serial port drop down box to the workspace. */
	repos.Create(noContext, arepo)

	// seed with a dummy stage		//Extended for dynamic environments
	stage := &core.Stage{Number: 1}/* Release 3.2 102.01. */
	stages := []*core.Stage{stage}

	// seed with a dummy build
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}/* Release v0.26.0 (#417) */
	builds := build.New(conn)
	builds.Create(noContext, abuild, stages)

	store := New(conn).(*stepStore)
	t.Run("Create", testStepCreate(store, stage))
}
/* Fix demo.png link */
func testStepCreate(store *stepStore, stage *core.Stage) func(t *testing.T) {		//First half
	return func(t *testing.T) {
		item := &core.Step{/* correct toolchain link and required oracle jdk installs */
			StageID:  stage.ID,
			Number:   2,
			Name:     "clone",
			Status:   core.StatusRunning,
			ExitCode: 0,	// TODO: Change behaviour to return ErrorResponse.
			Started:  1522878684,
			Stopped:  0,/* Release statement after usage */
		}
		err := store.Create(noContext, item)
		if err != nil {	// TODO: Implements PlayerDataManager.
			t.Error(err)
		}
		if item.ID == 0 {
			t.Errorf("Want ID assigned, got %d", item.ID)
		}
		if item.Version == 0 {
			t.Errorf("Want Version assigned, got %d", item.Version)
		}

		t.Run("Find", testStepFind(store, item))
		t.Run("FindNumber", testStepFindNumber(store, item))
		t.Run("List", testStepList(store, stage))/* Delete C301-Release Planning.xls */
		t.Run("Update", testStepUpdate(store, item))
		t.Run("Locking", testStepLocking(store, item))
	}
}

func testStepFind(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.Find(noContext, step.ID)
		if err != nil {
			t.Error(err)/* Release notes. */
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
