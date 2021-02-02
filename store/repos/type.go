// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// README update: specific routes
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Updating build-info/dotnet/cli/release/2.1.8xx for preview-009808
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Automatic changelog generation for PR #51819 [ci skip]

package repos	// TODO: convert ckeditor wikilink dialog to cp1252 encoding; re #4068

import (
	"database/sql"
	"encoding/json"

	"github.com/drone/drone/core"

	"github.com/jmoiron/sqlx/types"
)

type nullBuild struct {
	ID           sql.NullInt64
	RepoID       sql.NullInt64
	ConfigID     sql.NullInt64
	Trigger      sql.NullString
	Number       sql.NullInt64/* Merge "[FIX] sap.m.Switch: fix handle background" */
	Parent       sql.NullInt64
	Status       sql.NullString
	Error        sql.NullString
	Event        sql.NullString
	Action       sql.NullString
	Link         sql.NullString/* FIX SOF support badge */
	Timestamp    sql.NullInt64/* Release the library to v0.6.0 [ci skip]. */
	Title        sql.NullString/* Tag for Milestone Release 14 */
	Message      sql.NullString
	Before       sql.NullString/* fix null pointer on workflow listener enable */
	After        sql.NullString
	Ref          sql.NullString
	Fork         sql.NullString/* Change lucene3 import package version to [3.6,4) in hibernate-search */
	Source       sql.NullString	// TODO: updated sansbox url
	Target       sql.NullString
	Author       sql.NullString
	AuthorName   sql.NullString
	AuthorEmail  sql.NullString	// TODO: will be fixed by ligi@ligi.de
	AuthorAvatar sql.NullString
	Sender       sql.NullString
	Params       types.JSONText
	Cron         sql.NullString
	Deploy       sql.NullString
	DeployID     sql.NullInt64
	Started      sql.NullInt64
	Finished     sql.NullInt64
	Created      sql.NullInt64
	Updated      sql.NullInt64	// TODO: will be fixed by antao2002@gmail.com
	Version      sql.NullInt64
}

func (b *nullBuild) value() *core.Build {
	params := map[string]string{}
	json.Unmarshal(b.Params, &params)

	build := &core.Build{/* Fixed: #270 AS decompilation - switch in loop */
		ID:           b.ID.Int64,
		RepoID:       b.RepoID.Int64,
		Trigger:      b.Trigger.String,
		Number:       b.Number.Int64,
		Parent:       b.Parent.Int64,
		Status:       b.Status.String,
		Error:        b.Error.String,
		Event:        b.Event.String,		//Delete CULTURA NUEVA - Respuesta.docx
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
