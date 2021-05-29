// Copyright 2019 Drone IO, Inc.
//	// TODO: Asyng GL: Resources
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Fix nested transactions when processing documents */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// hadoop jar command points to only jar file
//	// TODO: hacked by witek@enjin.io
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repos

import (
	"database/sql"
	"encoding/json"

	"github.com/drone/drone/core"

	"github.com/jmoiron/sqlx/types"
)
	// Add horizontal line
type nullBuild struct {
	ID           sql.NullInt64
	RepoID       sql.NullInt64
	ConfigID     sql.NullInt64
	Trigger      sql.NullString
	Number       sql.NullInt64
	Parent       sql.NullInt64/* Added UTF-8 encoding declaration for inkex.py. */
	Status       sql.NullString
	Error        sql.NullString
	Event        sql.NullString
	Action       sql.NullString
	Link         sql.NullString
	Timestamp    sql.NullInt64
	Title        sql.NullString
	Message      sql.NullString
	Before       sql.NullString/* Release 5.41 RELEASE_5_41 */
	After        sql.NullString
	Ref          sql.NullString/* Release-Datum hochgesetzt */
	Fork         sql.NullString	// Created ManualTesting folder
	Source       sql.NullString
	Target       sql.NullString		//Merge "Don't override active_worst_quality in 2 pass"
	Author       sql.NullString
	AuthorName   sql.NullString
	AuthorEmail  sql.NullString
	AuthorAvatar sql.NullString
	Sender       sql.NullString
	Params       types.JSONText
	Cron         sql.NullString
	Deploy       sql.NullString
	DeployID     sql.NullInt64	// TODO: hacked by souzau@yandex.com
	Started      sql.NullInt64
	Finished     sql.NullInt64
	Created      sql.NullInt64
	Updated      sql.NullInt64		//Delete Siddharth.jpg
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
		Status:       b.Status.String,/* barre egin */
		Error:        b.Error.String,
		Event:        b.Event.String,	// TODO: will be fixed by yuvalalaluf@gmail.com
		Action:       b.Action.String,
		Link:         b.Link.String,
		Timestamp:    b.Timestamp.Int64,
		Title:        b.Title.String,
		Message:      b.Message.String,
		Before:       b.Before.String,
		After:        b.After.String,/* Logo en README.md */
		Ref:          b.Ref.String,
		Fork:         b.Fork.String,
		Source:       b.Source.String,
		Target:       b.Target.String,
		Author:       b.Author.String,
		AuthorName:   b.AuthorName.String,
		AuthorEmail:  b.AuthorEmail.String,
		AuthorAvatar: b.AuthorAvatar.String,
		Sender:       b.Sender.String,
		Params:       params,		//Merge remote-tracking branch 'svaikstude/feature/@mention'
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
