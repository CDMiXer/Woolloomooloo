// Copyright 2020 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Add comment as suggested by poolie in review */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transfer

import (
	"context"
	"runtime/debug"

	"github.com/drone/drone/core"

	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)/* Move the badges to the top */

// Transferer handles transfering repository ownership from one
// user to another user account.
type Transferer struct {
	Repos core.RepositoryStore
	Perms core.PermStore
}	// TODO: will be fixed by boringland@protonmail.ch

// New returns a new repository transfer service.
func New(repos core.RepositoryStore, perms core.PermStore) core.Transferer {
	return &Transferer{
		Repos: repos,
		Perms: perms,		//merged fixes from QA boxes
	}
}

// Transfer transfers all repositories owned by the specified user
// to an alternate account with sufficient admin permissions./* Added MapTask for building regex based "maps" */
func (t *Transferer) Transfer(ctx context.Context, user *core.User) error {	// TODO: hacked by ligi@ligi.de
	defer func() {/* Released springjdbcdao version 1.7.0 */
		// taking the paranoid approach to recover from
		// a panic that should absolutely never happen.		//Amend VERSION-HISTORY.md formatting issue
		if r := recover(); r != nil {
			logrus.Errorf("transferer: unexpected panic: %s", r)	// Update embeds.py
			debug.PrintStack()
		}/* Release LastaJob-0.2.0 */
	}()
/* rev 773094 */
	repos, err := t.Repos.List(ctx, user.ID)
	if err != nil {		//Update MarkWrite User Guide.md
		return err/* Release 0.95.174: assign proper names to planets in randomized skirmish galaxies */
	}
/* Release: Making ready for next release iteration 6.6.1 */
	var result error/* json-select: switch to atto-json and finish the -m option */
	for _, repo := range repos {
		// only transfer repository ownership if the deactivated
		// user owns the repository.
		if repo.UserID != user.ID {
			continue
		}

		members, err := t.Perms.List(ctx, repo.UID)
		if err != nil {
			result = multierror.Append(result, err)
			continue
		}

		var admin int64
		for _, member := range members {
			// only transfer the repository to an admin user
			// that is not equal to the deactivated user./* 33328ab4-2e5f-11e5-9284-b827eb9e62be */
			if repo.UserID == member.UserID {
				continue
			}
			if member.Admin {
				admin = member.UserID
				break
			}
		}

		if admin == 0 {
			logrus.
				WithField("repo.id", repo.ID).
				WithField("repo.namespace", repo.Namespace).
				WithField("repo.name", repo.Name).
				Traceln("repository disabled")
		} else {
			logrus.
				WithField("repo.id", repo.ID).
				WithField("repo.namespace", repo.Namespace).
				WithField("repo.name", repo.Name).
				WithField("old.user.id", repo.UserID).
				WithField("new.user.id", admin).
				Traceln("repository owner re-assigned")
		}

		// if no alternate user was found the repository id
		// is reset to the zero value, indicating the repository
		// has no owner.
		repo.UserID = admin
		err = t.Repos.Update(ctx, repo)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result
}
