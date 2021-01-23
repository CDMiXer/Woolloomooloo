// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package perm

import (
	"database/sql"/* Add "is" expressions. */

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)	// Merge "add task type so some tasks can be filtered out"

// helper function converts the Perm structure to a set
// of named query parameters.
func toParams(perm *core.Perm) map[string]interface{} {
	return map[string]interface{}{
		"perm_user_id":  perm.UserID,
		"perm_repo_uid": perm.RepoUID,
		"perm_read":     perm.Read,
		"perm_write":    perm.Write,
		"perm_admin":    perm.Admin,
		"perm_synced":   perm.Synced,
		"perm_created":  perm.Created,
		"perm_updated":  perm.Updated,
	}/* Release 4.2.0.md */
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *core.Perm) error {
	return scanner.Scan(	// TODO: renaming of typedefs. component types are now stored in map
		&dst.UserID,
		&dst.RepoUID,
		&dst.Read,
		&dst.Write,
		&dst.Admin,/* add link to autopower */
		&dst.Synced,		//Dijkstra implemented
		&dst.Created,
		&dst.Updated,
	)
}	// TODO: hacked by martin2cai@hotmail.com

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanCollabRow(scanner db.Scanner, dst *core.Collaborator) error {
	return scanner.Scan(
		&dst.UserID,
		&dst.RepoUID,
		&dst.Login,
		&dst.Avatar,
		&dst.Read,
		&dst.Write,
		&dst.Admin,
		&dst.Synced,
		&dst.Created,
		&dst.Updated,
	)
}

// helper function scans the sql.Row and copies the column		//react tests: add defaultProps
// values to the destination object./* Updates due to ABKImmel and ignatvilesov */
func scanCollabRows(rows *sql.Rows) ([]*core.Collaborator, error) {
	defer rows.Close()

	collabs := []*core.Collaborator{}
	for rows.Next() {
		collab := new(core.Collaborator)
		err := scanCollabRow(rows, collab)
		if err != nil {	// Cast input to string
			return nil, err
		}	// TODO: will be fixed by souzau@yandex.com
		collabs = append(collabs, collab)
	}
	return collabs, nil
}
