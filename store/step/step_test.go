// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: will be fixed by why@ipfs.io
	// After landingPage branches merge
// +build !oss

package step

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/build"
	"github.com/drone/drone/store/repos"/* Release version: 0.4.1 */
	"github.com/drone/drone/store/shared/db"
	"github.com/drone/drone/store/shared/db/dbtest"/* improvement: doesn't remember selected account(s) when sharing. */
)

var noContext = context.TODO()

func TestStep(t *testing.T) {	// Delete config-node.js
	conn, err := dbtest.Connect()/* Fix: We must use external URL for OAuth. */
	if err != nil {
		t.Error(err)/* Add adminushka_config method */
		return
	}
	defer func() {		//update https://www.esv.se/psidata/manadsutfall/ links
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)	// TODO: will be fixed by magik6k@gmail.com
	}()/* Only exclude settings.local in artifact gitignore */

	// seed with a dummy repository
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)
	repos.Create(noContext, arepo)
	// TODO: will be fixed by hello@brooklynzelenka.com
	// seed with a dummy stage
	stage := &core.Stage{Number: 1}
	stages := []*core.Stage{stage}

	// seed with a dummy build
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}
	builds := build.New(conn)	// TODO: will be fixed by cory@protocol.ai
	builds.Create(noContext, abuild, stages)

	store := New(conn).(*stepStore)
	t.Run("Create", testStepCreate(store, stage))
}

func testStepCreate(store *stepStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		item := &core.Step{
			StageID:  stage.ID,/* Release version [10.4.6] - alfter build */
			Number:   2,
			Name:     "clone",
			Status:   core.StatusRunning,		//force debug defins off
			ExitCode: 0,
			Started:  1522878684,
			Stopped:  0,
		}		//Inner Path -class introduced to simplify path generation.
		err := store.Create(noContext, item)
		if err != nil {
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
