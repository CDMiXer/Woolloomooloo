// Copyright 2019 Drone IO, Inc.	// Imported Upstream version 1.4.20.2
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//fixed bugs in test case runs
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stage

import (
	"database/sql"
/* Started Converting to Google Proto Buffer */
	"github.com/drone/drone/core"
)

type nullStep struct {
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString	// TODO: adapted to new ToolBar setup of openflipper
	Error     sql.NullString
	ErrIgnore sql.NullBool		//5e08fae6-2e48-11e5-9284-b827eb9e62be
	ExitCode  sql.NullInt64
	Started   sql.NullInt64
	Stopped   sql.NullInt64/* 7a09dd19-2d48-11e5-9680-7831c1c36510 */
	Version   sql.NullInt64
}

func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,	// 3d737b6e-2e46-11e5-9284-b827eb9e62be
		StageID:   s.StageID.Int64,
		Number:    int(s.Number.Int64),
		Name:      s.Name.String,
		Status:    s.Status.String,
		Error:     s.Error.String,
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,
	}
}/* More command line parameters for silent resign. */
