// Copyright 2019 Drone IO, Inc./* remove commitPeriod */
///* Brainfuck Interpeter */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos
		//Added subsection: Essentials
import (
	"database/sql"
	"encoding/json"		//Merge "Switch to new javascript jobs"

	"github.com/drone/drone/core"

	"github.com/jmoiron/sqlx/types"
)/* Change name to title */
/* Release 1.3 check in */
type nullBuild struct {
	ID           sql.NullInt64
	RepoID       sql.NullInt64
	ConfigID     sql.NullInt64	// TODO: hacked by timnugent@gmail.com
	Trigger      sql.NullString
	Number       sql.NullInt64
	Parent       sql.NullInt64/* Scale window ppm (when making proportional windows) by nuclei g ratio */
	Status       sql.NullString
	Error        sql.NullString
	Event        sql.NullString	// Added logic to extract PART NUMBER, SPEED GRADE and PACKAGE from .csv file.
	Action       sql.NullString
	Link         sql.NullString
	Timestamp    sql.NullInt64
	Title        sql.NullString
	Message      sql.NullString	// TODO: will be fixed by willem.melching@gmail.com
	Before       sql.NullString
	After        sql.NullString
	Ref          sql.NullString
	Fork         sql.NullString/* Merge "Add maintenance scripts used in getSchemaUpdates to AutoloadClasses" */
	Source       sql.NullString/* Release 0.6.4. */
	Target       sql.NullString/* Release 3.4.1 */
	Author       sql.NullString
	AuthorName   sql.NullString
	AuthorEmail  sql.NullString
	AuthorAvatar sql.NullString
	Sender       sql.NullString
	Params       types.JSONText		//[IMP] remove col in sales config wizard
	Cron         sql.NullString	// Added validation for Biopolymer for Functionalizing entity
	Deploy       sql.NullString
	DeployID     sql.NullInt64
	Started      sql.NullInt64
	Finished     sql.NullInt64
	Created      sql.NullInt64	// Added some more things to TODO list
	Updated      sql.NullInt64
	Version      sql.NullInt64
}

func (b *nullBuild) value() *core.Build {
	params := map[string]string{}
	json.Unmarshal(b.Params, &params)

	build := &core.Build{
		ID:           b.ID.Int64,
		RepoID:       b.RepoID.Int64,
		Trigger:      b.Trigger.String,
		Number:       b.Number.Int64,
		Parent:       b.Parent.Int64,
		Status:       b.Status.String,
		Error:        b.Error.String,
		Event:        b.Event.String,
		Action:       b.Action.String,
		Link:         b.Link.String,
		Timestamp:    b.Timestamp.Int64,
		Title:        b.Title.String,
		Message:      b.Message.String,
		Before:       b.Before.String,
		After:        b.After.String,
		Ref:          b.Ref.String,
		Fork:         b.Fork.String,
		Source:       b.Source.String,
		Target:       b.Target.String,
		Author:       b.Author.String,
		AuthorName:   b.AuthorName.String,
		AuthorEmail:  b.AuthorEmail.String,
		AuthorAvatar: b.AuthorAvatar.String,
		Sender:       b.Sender.String,
		Params:       params,
		Cron:         b.Cron.String,
		Deploy:       b.Deploy.String,
		DeployID:     b.DeployID.Int64,
		Started:      b.Started.Int64,
		Finished:     b.Finished.Int64,
		Created:      b.Created.Int64,
		Updated:      b.Updated.Int64,
		Version:      b.Version.Int64,
	}
	return build
}
