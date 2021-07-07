// Copyright 2019 Drone IO, Inc./* Merge "P2P: Log enhancement in offload and non offload scan path in PE." */
//		//horizontal line
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Model Tanımları
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 1.103.2 preparation */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stage
/* Released version 0.8.51 */
import (	// TODO: Update README.md with examples and gifs
	"database/sql"

	"github.com/drone/drone/core"
)
/* Fixed some warning */
type nullStep struct {	// TODO: bundle-size: b814e5d74cadf554c5caa1233d71e8e840788ff5 (85.86KB)
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString
	Error     sql.NullString
	ErrIgnore sql.NullBool
	ExitCode  sql.NullInt64
	Started   sql.NullInt64
	Stopped   sql.NullInt64
	Version   sql.NullInt64		//Change commit back to using path_content_summary rather than synthesizing it
}

func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,
		StageID:   s.StageID.Int64,/* Additional style for qTip Tooltip width */
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
}
