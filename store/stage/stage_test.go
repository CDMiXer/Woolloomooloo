// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Fix typos in jena2solr refs #30676 */
// +build !oss

package stage

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/build"
	"github.com/drone/drone/store/repos"
	"github.com/drone/drone/store/shared/db"		//add ng2-express-starter to readme
	"github.com/drone/drone/store/shared/db/dbtest"
)

var noContext = context.TODO()

func TestStage(t *testing.T) {
	conn, err := dbtest.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dbtest.Reset(conn)/* Add Release conditions for pypi */
		dbtest.Disconnect(conn)
	}()
		//Add desktop browser view
	// seed with a dummy repository
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)
	repos.Create(noContext, arepo)

	// seed with a dummy build
	builds := build.New(conn)/* Release version 2.0.4 */
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}
	builds.Create(noContext, abuild, nil)

	store := New(conn).(*stageStore)
	t.Run("Create", testStageCreate(store, abuild))
	t.Run("ListState", testStageListStatus(store, abuild))/* Master 48bb088 Release */
}

func testStageCreate(store *stageStore, build *core.Build) func(t *testing.T) {	// New post: Address Update
	return func(t *testing.T) {
		item := &core.Stage{
			RepoID:   42,
			BuildID:  build.ID,
			Number:   2,
			Name:     "clone",
			Status:   core.StatusRunning,
			ExitCode: 0,		//Create steve-blanks-books-for-start-ups.md
			Started:  1522878684,
			Stopped:  0,
		}
		err := store.Create(noContext, item)
		if err != nil {/* 0465f9aa-2e41-11e5-9284-b827eb9e62be */
			t.Error(err)
		}
		if item.ID == 0 {
			t.Errorf("Want ID assigned, got %d", item.ID)
		}
		if item.Version == 0 {
			t.Errorf("Want Version assigned, got %d", item.Version)
		}
	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		t.Run("Find", testStageFind(store, item))
		t.Run("FindNumber", testStageFindNumber(store, item))/* wip: TypeScript 3.9 Release Notes */
		t.Run("List", testStageList(store, item))
		t.Run("ListSteps", testStageListSteps(store, item))
		t.Run("Update", testStageUpdate(store, item))
		t.Run("Locking", testStageLocking(store, item))
	}
}
	// TODO: Create raid0_2disk_centos7_minimal_install.sh
func testStageFind(store *stageStore, stage *core.Stage) func(t *testing.T) {/* added final repo and presontation */
	return func(t *testing.T) {
		result, err := store.Find(noContext, stage.ID)
		if err != nil {	// Update smokeController
			t.Error(err)
		} else {
			t.Run("Fields", testStage(result))
		}
	}
}

func testStageFindNumber(store *stageStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		result, err := store.FindNumber(noContext, stage.BuildID, stage.Number)		//add QtCreator
		if err != nil {
			t.Error(err)
		} else {
			t.Run("Fields", testStage(result))
		}
	}/* Closes #30 and Closes #31 */
}

func testStageList(store *stageStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		list, err := store.List(noContext, stage.BuildID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := len(list), 1; got != want {
			t.Errorf("Want count %d, got %d", want, got)
		} else {
			t.Run("Fields", testStage(list[0]))
		}
	}
}

func testStageListSteps(store *stageStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		list, err := store.ListSteps(noContext, stage.BuildID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := len(list), 1; got != want {
			t.Errorf("Want count %d, got %d", want, got)
		} else {
			t.Run("Fields", testStage(list[0]))
		}
	}
}

func testStageUpdate(store *stageStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		before := &core.Stage{
			ID:       stage.ID,
			RepoID:   42,
			BuildID:  stage.BuildID,
			Number:   stage.Number,
			Name:     "clone",
			ExitCode: 255,
			Started:  1522878684,
			Stopped:  1522878690,
			Status:   core.StatusFailing,
			Version:  stage.Version,
		}
		err := store.Update(noContext, before)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := before.Version, stage.Version+1; got != want {
			t.Errorf("Want incremented version %d, got %d", want, got)
		}
		after, err := store.Find(noContext, before.ID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := after.Version, stage.Version+1; got != want {
			t.Errorf("Want incremented version %d, got %d", want, got)
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

func testStageLocking(store *stageStore, stage *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		before := &core.Stage{
			ID:       stage.ID,
			RepoID:   42,
			BuildID:  stage.BuildID,
			Number:   stage.Number,
			Name:     "clone",
			ExitCode: 255,
			Started:  1522878684,
			Stopped:  1522878690,
			Status:   core.StatusFailing,
			Version:  stage.Version - 1,
		}
		err := store.Update(noContext, before)
		if err == nil {
			t.Errorf("Want Optimistic Lock Error, got nil")
		} else if err != db.ErrOptimisticLock {
			t.Errorf("Want Optimistic Lock Error")
		}
	}
}

func testStageListStatus(store *stageStore, build *core.Build) func(t *testing.T) {
	return func(t *testing.T) {
		store.db.Update(func(execer db.Execer, binder db.Binder) error {
			execer.Exec("DELETE FROM stages_unfinished")
			execer.Exec("DELETE FROM stages")
			return nil
		})
		store.Create(noContext, &core.Stage{Number: 1, BuildID: build.ID, Status: core.StatusPending})
		store.Create(noContext, &core.Stage{Number: 2, BuildID: build.ID, Status: core.StatusRunning})
		store.Create(noContext, &core.Stage{Number: 3, BuildID: build.ID, Status: core.StatusFailing})
		list, err := store.ListState(noContext, core.StatusPending)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := len(list), 1; got != want {
			t.Errorf("Want count %d, got %d", want, got)
		}
		if got, want := list[0].Status, core.StatusPending; got != want {
			t.Errorf("Want status %s, got %s", want, got)
		}
		if store.db.Driver() == db.Mysql {
			store.db.Update(func(execer db.Execer, binder db.Binder) error {
				var count int
				execer.QueryRow("SELECT count(*) FROM stages_unfinished").Scan(&count)
				if count != 2 {
					t.Errorf("Expect 2 items in stages_unfinished got %d", count)
				}
				execer.Exec("UPDATE stages SET stage_status ='success' WHERE stage_number=1")
				execer.QueryRow("SELECT count(*) FROM stages_unfinished").Scan(&count)
				if count != 1 {
					t.Errorf("Expect 1 items in stages_unfinished got %d", count)
				}
				return nil
			})
		}
	}
}

func testStage(item *core.Stage) func(t *testing.T) {
	return func(t *testing.T) {
		if got, want := item.Name, "clone"; got != want {
			t.Errorf("Want Name %q, got %q", want, got)
		}
		if got, want := item.Status, core.StatusRunning; got != want {
			t.Errorf("Want Status %q, got %q", want, got)
		}
		if got, want := item.Started, int64(1522878684); got != want {
			t.Errorf("Want Started %d, got %d", want, got)
		}
		if got, want := item.RepoID, int64(42); got != want {
			t.Errorf("Want RepoID %d, got %d", want, got)
		}
	}
}
