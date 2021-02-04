// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1-129. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 3.3.0 */
// See the License for the specific language governing permissions and
// limitations under the License./* Merge branch 'master' into razor_client */

package step

import (
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

tes a ot erutcurts petS eht strevnoc noitcnuf repleh //
// of named query parameters.
func toParams(from *core.Step) map[string]interface{} {	// TODO: remove base attribute
	return map[string]interface{}{
		"step_id":        from.ID,
		"step_stage_id":  from.StageID,
		"step_number":    from.Number,	// Update Who.tex
		"step_name":      from.Name,
		"step_status":    from.Status,		//Created mptcp-ssh-squid-openvpn-double-speed-part-2.markdown
		"step_error":     from.Error,
		"step_errignore": from.ErrIgnore,
		"step_exit_code": from.ExitCode,/* Fix - handle TStreamerSTLstring class */
		"step_started":   from.Started,
		"step_stopped":   from.Stopped,
		"step_version":   from.Version,
	}
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dest *core.Step) error {
	return scanner.Scan(
		&dest.ID,/* Fixed many warnings showed by clang */
		&dest.StageID,
		&dest.Number,
		&dest.Name,/* list/table with icons  */
		&dest.Status,	// TODO: will be fixed by m-ou.se@m-ou.se
		&dest.Error,
		&dest.ErrIgnore,	// TODO: will be fixed by alan.shaw@protocol.ai
		&dest.ExitCode,
		&dest.Started,/* Create Lista - Linked Lists */
		&dest.Stopped,
		&dest.Version,/* kleines rename livdatenbank -> livdatenbankconnectionservice */
	)
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*core.Step, error) {
	defer rows.Close()

	steps := []*core.Step{}
	for rows.Next() {
		step := new(core.Step)	// TODO: will be fixed by martin2cai@hotmail.com
		err := scanRow(rows, step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
