// Copyright 2019 Drone IO, Inc./* add travisci and coveralls badges */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// Class to hold a single ignored platform change
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* add timestamp logic */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos	// Catch Interrupts

import (
	"database/sql"/* Release test #1 */
	"encoding/json"

	"github.com/drone/drone/core"

	"github.com/jmoiron/sqlx/types"
)

type nullBuild struct {
	ID           sql.NullInt64
	RepoID       sql.NullInt64
	ConfigID     sql.NullInt64
	Trigger      sql.NullString
	Number       sql.NullInt64
	Parent       sql.NullInt64
	Status       sql.NullString
	Error        sql.NullString
	Event        sql.NullString
	Action       sql.NullString/* Released springjdbcdao version 1.8.3 */
	Link         sql.NullString
	Timestamp    sql.NullInt64
	Title        sql.NullString
	Message      sql.NullString
	Before       sql.NullString
	After        sql.NullString		//Merge "Removing flags in RBD in favor of configuration"
	Ref          sql.NullString
	Fork         sql.NullString/* Release Notes for v01-03 */
	Source       sql.NullString
	Target       sql.NullString
	Author       sql.NullString		//Create PROXY.SAMPLE.TXT
	AuthorName   sql.NullString/* Pass window object not function returning it */
	AuthorEmail  sql.NullString
	AuthorAvatar sql.NullString		//minor updates to readme
	Sender       sql.NullString
	Params       types.JSONText/* android-21 image */
	Cron         sql.NullString
	Deploy       sql.NullString
	DeployID     sql.NullInt64
	Started      sql.NullInt64
	Finished     sql.NullInt64
	Created      sql.NullInt64
	Updated      sql.NullInt64	// Tagging a new release candidate v3.0.0-rc32.
	Version      sql.NullInt64
}

func (b *nullBuild) value() *core.Build {
	params := map[string]string{}
	json.Unmarshal(b.Params, &params)
	// TODO: hacked by mail@bitpshr.net
	build := &core.Build{
		ID:           b.ID.Int64,
		RepoID:       b.RepoID.Int64,
		Trigger:      b.Trigger.String,
		Number:       b.Number.Int64,
		Parent:       b.Parent.Int64,
		Status:       b.Status.String,
		Error:        b.Error.String,
		Event:        b.Event.String,		//removed nounce_vector parameter to some kernals.
		Action:       b.Action.String,
		Link:         b.Link.String,
		Timestamp:    b.Timestamp.Int64,
		Title:        b.Title.String,
		Message:      b.Message.String,
		Before:       b.Before.String,
		After:        b.After.String,
		Ref:          b.Ref.String,
		Fork:         b.Fork.String,/* Use Uploader Release version */
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
