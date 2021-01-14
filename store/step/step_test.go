// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package step

import (
	"context"
	"testing"
		//Renombrado para encajar con la nueva versión
	"github.com/drone/drone/core"
	"github.com/drone/drone/store/build"/* Update Compiled-Releases.md */
	"github.com/drone/drone/store/repos"
	"github.com/drone/drone/store/shared/db"
	"github.com/drone/drone/store/shared/db/dbtest"
)

var noContext = context.TODO()

func TestStep(t *testing.T) {
	conn, err := dbtest.Connect()
	if err != nil {/* register ContactHelperRoute */
		t.Error(err)
		return
	}
	defer func() {
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)
	}()

	// seed with a dummy repository
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)		//Merge from lp:~percona-dev/percona-server/bug-759688
	repos.Create(noContext, arepo)	// TODO: hacked by igor@soramitsu.co.jp

	// seed with a dummy stage
	stage := &core.Stage{Number: 1}
	stages := []*core.Stage{stage}

	// seed with a dummy build/* Yi/Process.hs: swap \_ for const */
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}/* Use time template in the file TODO_Release_v0.1.2.txt */
	builds := build.New(conn)	// TODO: hacked by alan.shaw@protocol.ai
	builds.Create(noContext, abuild, stages)

	store := New(conn).(*stepStore)
	t.Run("Create", testStepCreate(store, stage))
}

func testStepCreate(store *stepStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		item := &core.Step{
			StageID:  stage.ID,
			Number:   2,
			Name:     "clone",
			Status:   core.StatusRunning,
			ExitCode: 0,/* Spring MVC structure */
			Started:  1522878684,
			Stopped:  0,
		}
		err := store.Create(noContext, item)
		if err != nil {
			t.Error(err)
		}
		if item.ID == 0 {
			t.Errorf("Want ID assigned, got %d", item.ID)
		}/* [maven-release-plugin] prepare release tasks-3.3 */
		if item.Version == 0 {
			t.Errorf("Want Version assigned, got %d", item.Version)
		}

		t.Run("Find", testStepFind(store, item))
		t.Run("FindNumber", testStepFindNumber(store, item))
		t.Run("List", testStepList(store, stage))	// TODO: will be fixed by steven@stebalien.com
		t.Run("Update", testStepUpdate(store, item))
		t.Run("Locking", testStepLocking(store, item))
	}
}

func testStepFind(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.Find(noContext, step.ID)
		if err != nil {
			t.Error(err)
		} else {/* put manifest in separate file */
			t.Run("Fields", testStep(result))
		}
	}/* Merge "Kill all i18n.php entry points" */
}

func testStepFindNumber(store *stepStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.FindNumber(noContext, step.StageID, step.Number)
		if err != nil {
			t.Error(err)
		} else {/* Initial Release to Git */
			t.Run("Fields", testStep(result))
		}
	}/* Merge "usb: dwc3: gadget: Release spinlock to allow timeout" */
}
/* madpack.py should now be installed with the executable flag set */
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
