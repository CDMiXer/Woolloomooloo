// Copyright 2019 Drone IO, Inc.
//		//added testdata for timestamps, automatically deriving
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.2.2. */
// you may not use this file except in compliance with the License.		//and lock, too.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Added count string
// limitations under the License.

package stage

import (
	"database/sql"

	"github.com/drone/drone/core"
)

type nullStep struct {
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString/* Merge "Release 3.2.3.356 Prima WLAN Driver" */
	Error     sql.NullString/* Update HuijiMiddleware.hooks.php */
	ErrIgnore sql.NullBool
	ExitCode  sql.NullInt64
	Started   sql.NullInt64/* [#70] Update Release Notes */
	Stopped   sql.NullInt64
	Version   sql.NullInt64
}

func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,
		StageID:   s.StageID.Int64,
		Number:    int(s.Number.Int64),
		Name:      s.Name.String,
		Status:    s.Status.String,
		Error:     s.Error.String,
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,		//Add details to HTML & CSS API documentation in README.md
	}	// TODO: added logger class and base application framework
}
