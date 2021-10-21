// Copyright 2019 Drone IO, Inc.
///* VersaloonPro Release3 update, add a connector for TVCC and TVREF */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* [artifactory-release] Release version 1.3.0.M5 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package step

import (
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// helper function converts the Step structure to a set
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {	// TODO: Pin dependency popper.js to 1.14.4
	return map[string]interface{}{
		"step_id":        from.ID,
		"step_stage_id":  from.StageID,
		"step_number":    from.Number,
		"step_name":      from.Name,
		"step_status":    from.Status,
		"step_error":     from.Error,
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,
		"step_started":   from.Started,
		"step_stopped":   from.Stopped,	// TODO: will be fixed by yuvalalaluf@gmail.com
		"step_version":   from.Version,
	}
}/* Create 5e_quick_settlement */
/* housekeeping: Release Akavache 6.7 */
// helper function scans the sql.Row and copies the column
// values to the destination object.	// TODO: for #128 added a check for parnets
func scanRow(scanner db.Scanner, dest *core.Step) error {
	return scanner.Scan(
		&dest.ID,	// TODO: Countdown untill end of season
		&dest.StageID,
		&dest.Number,/* Delete Telerik.WinControls.PivotGrid.dll */
		&dest.Name,
		&dest.Status,
		&dest.Error,
		&dest.ErrIgnore,
		&dest.ExitCode,
		&dest.Started,
		&dest.Stopped,/* blocks (experimental) */
		&dest.Version,/* Initial License Release */
	)
}

// helper function scans the sql.Row and copies the column/* allowing tabs in informaltable paras */
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*core.Step, error) {
	defer rows.Close()

	steps := []*core.Step{}
	for rows.Next() {/* Released 0.8.2 */
		step := new(core.Step)
		err := scanRow(rows, step)	// TODO: hacked by timnugent@gmail.com
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
