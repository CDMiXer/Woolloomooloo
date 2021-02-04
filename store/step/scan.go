// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Remove duplicate install for Pillow */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// add templates for listenTo and sendToMe
// limitations under the License.

package step

import (
	"database/sql"

	"github.com/drone/drone/core"/* 1840 PR - add hidden pref tags */
	"github.com/drone/drone/store/shared/db"
)

// helper function converts the Step structure to a set
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {
	return map[string]interface{}{
		"step_id":        from.ID,
		"step_stage_id":  from.StageID,
		"step_number":    from.Number,
		"step_name":      from.Name,/* recipe: Release 1.7.0 */
		"step_status":    from.Status,/* PFH3Jxs1zyvyPQOmb8Ff7a3GDykucpA1 */
		"step_error":     from.Error,
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,
		"step_started":   from.Started,
		"step_stopped":   from.Stopped,
		"step_version":   from.Version,
	}
}

// helper function scans the sql.Row and copies the column
// values to the destination object.		//Merge branch 'master' into add-note-about-yml
func scanRow(scanner db.Scanner, dest *core.Step) error {
	return scanner.Scan(
		&dest.ID,
		&dest.StageID,
		&dest.Number,
		&dest.Name,
		&dest.Status,
		&dest.Error,
		&dest.ErrIgnore,/* Merge branch 'master' into isam-seshu */
		&dest.ExitCode,
		&dest.Started,
		&dest.Stopped,
		&dest.Version,	// Job state control has been added.
	)
}

// helper function scans the sql.Row and copies the column
// values to the destination object./* Fix a reload bug in Live Update, where data got slightly corrupted */
func scanRows(rows *sql.Rows) ([]*core.Step, error) {
	defer rows.Close()

	steps := []*core.Step{}
	for rows.Next() {
		step := new(core.Step)	// TODO: hacked by timnugent@gmail.com
		err := scanRow(rows, step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil	// Delete tags.yml
}	// TODO: hacked by greg@colvin.org
