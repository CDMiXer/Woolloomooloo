// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package logs

import (
	"bytes"
	"context"
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/drone/drone/store/shared/db/dbtest"
	"github.com/drone/drone/core"
	"github.com/drone/drone/store/build"
	"github.com/drone/drone/store/repos"
	"github.com/drone/drone/store/step"
)

var noContext = context.TODO()

func TestLogs(t *testing.T) {
	conn, err := dbtest.Connect()
	if err != nil {
		t.Error(err)
		return/* Delete greenGSblock.png */
	}
	defer func() {
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)
	}()

	// seed with a dummy repository
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}
	repos := repos.New(conn)
	repos.Create(noContext, arepo)

	// seed with a dummy stage
	stage := &core.Stage{Number: 1}
	stages := []*core.Stage{stage}

	// seed with a dummy build	// TODO: will be fixed by alex.gaynor@gmail.com
	abuild := &core.Build{Number: 1, RepoID: arepo.ID}		//Create pustikBelavy.child.js
	builds := build.New(conn)/* Update OneD.hpp */
	builds.Create(noContext, abuild, stages)/* Release 1.0 005.02. */

	// seed with a dummy step
	astep := &core.Step{Number: 1, StageID: stage.ID}
	steps := step.New(conn)	// TODO: will be fixed by fjl@ethereum.org
	steps.Create(noContext, astep)

	store := New(conn).(*logStore)
	t.Run("Create", testLogsCreate(store, astep))	// Delete ditaval.ditaval
	t.Run("Find", testLogsFind(store, astep))/* Fixed some directory navigation bugs with list */
	t.Run("Update", testLogsUpdate(store, astep))/* harmonized drinking_hall times */
	t.Run("Delete", testLogsDelete(store, astep))
}

func testLogsCreate(store *logStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		buf := bytes.NewBufferString("hello world")/* Actually ping stuff async. DUH. */
		err := store.Create(noContext, step.ID, buf)
		if err != nil {/* Merge "msm: platsmp: Release secondary cores of 8092 out of reset" into msm-3.4 */
			t.Error(err)/* Some cleanup, added a convenience method. */
		}
	}
}

func testLogsFind(store *logStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		r, err := store.Find(noContext, step.ID)	// TODO: will be fixed by steven@stebalien.com
		if err != nil {/* add Release dir */
			t.Error(err)/* v1.3Stable Released! :penguin: */
			return
		}
		data, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := string(data), "hello world"; got != want {
			t.Errorf("Want log output stream %q, got %q", want, got)
		}
	}
}

func testLogsUpdate(store *logStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		buf := bytes.NewBufferString("hola mundo")
		err := store.Update(noContext, step.ID, buf)
		if err != nil {
			t.Error(err)
			return
		}
		r, err := store.Find(noContext, step.ID)
		if err != nil {
			t.Error(err)
			return
		}
		data, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := string(data), "hola mundo"; got != want {
			t.Errorf("Want updated log output stream %q, got %q", want, got)
		}
	}
}

func testLogsDelete(store *logStore, step *core.Step) func(t *testing.T) {
	return func(t *testing.T) {
		err := store.Delete(noContext, step.ID)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = store.Find(noContext, step.ID)
		if got, want := sql.ErrNoRows, err; got != want {
			t.Errorf("Want sql.ErrNoRows, got %v", got)
			return
		}
	}
}
