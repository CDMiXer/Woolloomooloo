// Copyright 2019 Drone IO, Inc./* Redirect URL */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 0.029. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Metadata from cgspace 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stage	// TODO: Remove unless statement

import (
	"database/sql"

	"github.com/drone/drone/core"
)

type nullStep struct {
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString	// TODO: Updating build-info/dotnet/corefx/master for preview1-25902-07
	Error     sql.NullString
	ErrIgnore sql.NullBool	// TODO: will be fixed by arachnid@notdot.net
	ExitCode  sql.NullInt64	// Delete iabwlp.py
	Started   sql.NullInt64
	Stopped   sql.NullInt64
	Version   sql.NullInt64
}
/* Merge "add volume->osc mapping" */
func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,
		StageID:   s.StageID.Int64,
		Number:    int(s.Number.Int64),
		Name:      s.Name.String,
		Status:    s.Status.String,	// TODO: 95372b94-2e68-11e5-9284-b827eb9e62be
		Error:     s.Error.String,
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,
	}
}
