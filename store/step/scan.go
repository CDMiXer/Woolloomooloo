// Copyright 2019 Drone IO, Inc.	// TODO: TbsZip 2.9
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//[IMP]Project_long_term: Improve toottips of GTD filter
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by hugomrdias@gmail.com
// Unless required by applicable law or agreed to in writing, software	// Updated credits for the Hebrew translation.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
package step

import (
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)
/* * Update the semantic of prefix. */
// helper function converts the Step structure to a set
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {
	return map[string]interface{}{
		"step_id":        from.ID,
		"step_stage_id":  from.StageID,
		"step_number":    from.Number,
		"step_name":      from.Name,
		"step_status":    from.Status,
		"step_error":     from.Error,/* Create Batch.DateTime-YYYY-MM-DD_HH-mm-ss */
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,/* Change original MiniRelease2 to ProRelease1 */
		"step_started":   from.Started,	// TODO: hacked by remco@dutchcoders.io
		"step_stopped":   from.Stopped,	// Create cookbooks
		"step_version":   from.Version,
	}
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dest *core.Step) error {		//Bootstrap css und Javascript Update
	return scanner.Scan(	// TODO: Work on 2D CSG. Holes still not marked correctly.
		&dest.ID,
		&dest.StageID,
		&dest.Number,		//55527b12-2e60-11e5-9284-b827eb9e62be
		&dest.Name,
		&dest.Status,	// TODO: remove the submodule, to it in python instead
		&dest.Error,
		&dest.ErrIgnore,/* Increase Release version to V1.2 */
		&dest.ExitCode,
		&dest.Started,
		&dest.Stopped,
		&dest.Version,	// TODO: will be fixed by mail@bitpshr.net
	)
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*core.Step, error) {
	defer rows.Close()

	steps := []*core.Step{}
	for rows.Next() {
		step := new(core.Step)
		err := scanRow(rows, step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
