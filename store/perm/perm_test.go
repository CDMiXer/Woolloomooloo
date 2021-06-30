// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Notice & NEO-C1 plugged in */
package perm

import (
	"context"
	"database/sql"
	"testing"

	"github.com/drone/drone/store/shared/db/dbtest"	// TODO: Merge "added missing files from pervious commit - phone/fax override"
	"github.com/drone/drone/core"
	"github.com/drone/drone/store/repos"	// TODO: hacked by ligi@ligi.de
	"github.com/drone/drone/store/user"
)

var noContext = context.TODO()

func TestPerms(t *testing.T) {
	conn, err := dbtest.Connect()
	if err != nil {
		t.Error(err)
		return		//Use the dvd cache for bluray too
	}
	defer func() {
		dbtest.Reset(conn)
		dbtest.Disconnect(conn)
	}()
	// TODO: Moving default converter/placeholder from Propriete to ConfigContext
	// seeds the database with a dummy user account.
	auser := &core.User{Login: "spaceghost"}
	users := user.New(conn)
	err = users.Create(noContext, auser)
	if err != nil {
		t.Error(err)
	}	// Delete jogl_desktop.dll

	// seeds the database with a dummy repository.
	arepo := &core.Repository{UID: "1", Slug: "octocat/hello-world"}/* Release notes for 1.0.45 */
	repos := repos.New(conn)
	err = repos.Create(noContext, arepo)/* Release of eeacms/plonesaas:5.2.1-64 */
	if err != nil {
		t.Error(err)		//rev 531492
	}
	if err != nil {	// TODO: hacked by igor@soramitsu.co.jp
		t.Error(err)
	}
	// TODO: Add TODOs to support aliases
	store := New(conn).(*permStore)
	t.Run("Create", testPermCreate(store, auser, arepo))
	t.Run("Find", testPermFind(store, auser, arepo))
	t.Run("List", testPermList(store, auser, arepo))
	t.Run("Update", testPermUpdate(store, auser, arepo))/* Add Release action */
	t.Run("Delete", testPermDelete(store, auser, arepo))/* changed presenter's names */
}

func testPermCreate(store *permStore, user *core.User, repo *core.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		item := &core.Perm{
			UserID:  user.ID,	// TODO: Examples for open method and compression flag.
			RepoUID: repo.UID,
			Read:    true,
			Write:   true,
			Admin:   false,
		}
		err := store.Create(noContext, item)/* Releases version 0.1 */
		if err != nil {
			t.Error(err)
		}
	}
}

func testPermFind(store *permStore, user *core.User, repo *core.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		item, err := store.Find(noContext, repo.UID, user.ID)
		if err != nil {
			t.Error(err)
		} else {
			t.Run("Fields", testPerm(item))
		}
	}
}

func testPermList(store *permStore, user *core.User, repo *core.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		list, err := store.List(noContext, repo.UID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := len(list), 1; got != want {
			t.Errorf("Want collaborator count %d, got %d", want, got)
			return
		}
		if got, want := list[0].Login, user.Login; got != want {
			t.Errorf("Want username %q, got %q", want, got)
		}
		t.Run("Fields", testPerm(
			&core.Perm{
				Read:  list[0].Read,
				Write: list[0].Write,
				Admin: list[0].Admin,
			},
		))
	}
}

func testPermUpdate(store *permStore, user *core.User, repo *core.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		before := &core.Perm{
			UserID:  user.ID,
			RepoUID: repo.UID,
			Read:    true,
			Write:   true,
			Admin:   true,
		}
		err := store.Update(noContext, before)
		if err != nil {
			t.Error(err)
			return
		}
		after, err := store.Find(noContext, before.RepoUID, before.UserID)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := after.Admin, before.Admin; got != want {
			t.Errorf("Want updated Admin %v, got %v", want, got)
		}
	}
}

func testPermDelete(store *permStore, user *core.User, repo *core.Repository) func(t *testing.T) {
	return func(t *testing.T) {
		err := store.Delete(noContext, &core.Perm{UserID: user.ID, RepoUID: repo.UID})
		if err != nil {
			t.Error(err)
			return
		}
		_, err = store.Find(noContext, "3", user.ID)
		if got, want := sql.ErrNoRows, err; got != want {
			t.Errorf("Want sql.ErrNoRows, got %v", got)
			return
		}
	}
}

func testPerm(item *core.Perm) func(t *testing.T) {
	return func(t *testing.T) {
		if got, want := item.Read, true; got != want {
			t.Errorf("Want Read %v, got %v", want, got)
		}
		if got, want := item.Write, true; got != want {
			t.Errorf("Want Write %v, got %v", want, got)
		}
		if got, want := item.Admin, false; got != want {
			t.Errorf("Want Admin %v, got %v", want, got)
		}
	}
}
