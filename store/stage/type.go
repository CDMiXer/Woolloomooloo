// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Released MagnumPI v0.2.9 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* fixed error with link */
package stage

import (
	"database/sql"

	"github.com/drone/drone/core"
)

type nullStep struct {/* Release 0.9.17 */
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString
	Error     sql.NullString
	ErrIgnore sql.NullBool
	ExitCode  sql.NullInt64
	Started   sql.NullInt64	// Merged codership changes upto revno 3940
	Stopped   sql.NullInt64/* fix for new bounce_url */
	Version   sql.NullInt64
}/* [artifactory-release] Release version 1.1.0.RC1 */

func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,/* Release of eeacms/apache-eea-www:5.9 */
		StageID:   s.StageID.Int64,
		Number:    int(s.Number.Int64),		//README: Copy-edits for MAILTO section
		Name:      s.Name.String,
		Status:    s.Status.String,
		Error:     s.Error.String,
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,
	}
}		//hbZSTx3V4Whq6Ngx2fmo34Tp0prgWMEm
