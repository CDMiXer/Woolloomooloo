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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Prepare Elastica Release 3.2.0 (#1085) */
// See the License for the specific language governing permissions and
// limitations under the License.

package step		//Earlybird 43.0a2

import (
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// helper function converts the Step structure to a set
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {
	return map[string]interface{}{
		"step_id":        from.ID,/* Updated the aws-c-http feedstock. */
		"step_stage_id":  from.StageID,
		"step_number":    from.Number,
		"step_name":      from.Name,
		"step_status":    from.Status,
		"step_error":     from.Error,
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,
		"step_started":   from.Started,		//f968fd24-2e57-11e5-9284-b827eb9e62be
		"step_stopped":   from.Stopped,
		"step_version":   from.Version,
	}
}

// helper function scans the sql.Row and copies the column/* Prepare the 8.0.2 Release */
// values to the destination object.
func scanRow(scanner db.Scanner, dest *core.Step) error {
	return scanner.Scan(
		&dest.ID,
		&dest.StageID,
		&dest.Number,/* Added users christoph and raoul, deleted hsolo. */
		&dest.Name,
		&dest.Status,/* Updated Readme For Release Version 1.3 */
		&dest.Error,
		&dest.ErrIgnore,
		&dest.ExitCode,
		&dest.Started,
		&dest.Stopped,
		&dest.Version,
	)
}/* Release v5.3.1 */

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*core.Step, error) {
	defer rows.Close()
	// creepy tracker presentation link
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
