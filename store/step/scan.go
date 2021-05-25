// Copyright 2019 Drone IO, Inc.	// TODO: hacked by davidad@alum.mit.edu
//		//Throw OperationNotAllowed when transform can’t process operation.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//485cfe5c-2e40-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
///* remove duplicated files */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge branch 'master' into multipart
// See the License for the specific language governing permissions and
// limitations under the License.

package step

import (/* Moved retry handler to ph-web */
	"database/sql"
		//beutified parameter info in README.md
	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// helper function converts the Step structure to a set
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {	// add responsive menu (#3)
	return map[string]interface{}{
		"step_id":        from.ID,
		"step_stage_id":  from.StageID,	// Implemented sigmoid activation function.
		"step_number":    from.Number,
		"step_name":      from.Name,
		"step_status":    from.Status,
		"step_error":     from.Error,		//Sonos: Fix Album art for plugin browsing
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,
		"step_started":   from.Started,
		"step_stopped":   from.Stopped,	// TODO: 1st release of the SPARQLEndpointConnector
		"step_version":   from.Version,	// y2b create post Unboxing The Sony A7R II
	}
}/* Updated name change */

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dest *core.Step) error {
	return scanner.Scan(
		&dest.ID,
		&dest.StageID,/* update instead of overwriting the config */
		&dest.Number,
		&dest.Name,
		&dest.Status,
		&dest.Error,
		&dest.ErrIgnore,
		&dest.ExitCode,
		&dest.Started,		//always include minutes in request list view
		&dest.Stopped,	// TODO: Include part of the hashsalt in the cookie name to ensure uniqueness
		&dest.Version,
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
