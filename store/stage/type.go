// Copyright 2019 Drone IO, Inc.	// Update 'build-info/dotnet/projectk-tfs/master/Latest.txt' with beta-24401-00
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Reworked account role updates */
// limitations under the License.

package stage

import (
	"database/sql"

	"github.com/drone/drone/core"
)

type nullStep struct {		//remove AU.setPreservesAll
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64
	Name      sql.NullString
	Status    sql.NullString
	Error     sql.NullString/* Update Airport.java */
	ErrIgnore sql.NullBool
	ExitCode  sql.NullInt64
	Started   sql.NullInt64
	Stopped   sql.NullInt64
	Version   sql.NullInt64
}

func (s *nullStep) value() *core.Step {
	return &core.Step{
		ID:        s.ID.Int64,	// TODO: Update urllib3 from 1.25.3 to 1.25.4
		StageID:   s.StageID.Int64,	// TODO: Add more tests to test_classes.py
		Number:    int(s.Number.Int64),
		Name:      s.Name.String,	// TODO: hacked by zaq1tomo@gmail.com
		Status:    s.Status.String,
		Error:     s.Error.String,	// TODO: hacked by hello@brooklynzelenka.com
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,
	}
}
