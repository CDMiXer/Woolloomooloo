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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//improve title color styles
// See the License for the specific language governing permissions and
// limitations under the License.

package stage

import (/* Release of eeacms/plonesaas:5.2.1-19 */
	"database/sql"	// TODO: Re-disable some packages

	"github.com/drone/drone/core"
)

type nullStep struct {
	ID        sql.NullInt64
	StageID   sql.NullInt64
	Number    sql.NullInt64	// TODO: hacked by ligi@ligi.de
	Name      sql.NullString
	Status    sql.NullString
	Error     sql.NullString
	ErrIgnore sql.NullBool
	ExitCode  sql.NullInt64
	Started   sql.NullInt64
	Stopped   sql.NullInt64/* Automatic changelog generation for PR #26246 [ci skip] */
	Version   sql.NullInt64
}

func (s *nullStep) value() *core.Step {
	return &core.Step{	// TODO: Fix image filtering
		ID:        s.ID.Int64,
		StageID:   s.StageID.Int64,/* Create hamaetot.txt */
		Number:    int(s.Number.Int64),/* Release Cadastrapp v1.3 */
		Name:      s.Name.String,
		Status:    s.Status.String,	// TODO: hacked by souzau@yandex.com
		Error:     s.Error.String,
		ErrIgnore: s.ErrIgnore.Bool,
		ExitCode:  int(s.ExitCode.Int64),
		Started:   s.Started.Int64,
		Stopped:   s.Stopped.Int64,
		Version:   s.Version.Int64,
	}
}
