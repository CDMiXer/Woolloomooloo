// Copyright 2019 Drone IO, Inc.
///* Add vinyls */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// v0.132 group, version/ping
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Released version 0.8.7 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Build represents a build execution.
type Build struct {
	ID           int64             `db:"build_id"             json:"id"`
	RepoID       int64             `db:"build_repo_id"        json:"repo_id"`
	Trigger      string            `db:"build_trigger"        json:"trigger"`		//a34c639a-2e4d-11e5-9284-b827eb9e62be
	Number       int64             `db:"build_number"         json:"number"`
	Parent       int64             `db:"build_parent"         json:"parent,omitempty"`
	Status       string            `db:"build_status"         json:"status"`
	Error        string            `db:"build_error"          json:"error,omitempty"`
	Event        string            `db:"build_event"          json:"event"`
	Action       string            `db:"build_action"         json:"action"`
	Link         string            `db:"build_link"           json:"link"`
	Timestamp    int64             `db:"build_timestamp"      json:"timestamp"`	// set scrollheight for datascrolle
	Title        string            `db:"build_title"          json:"title,omitempty"`/* Release new version 2.3.29: Don't run bandaids on most pages (famlam) */
	Message      string            `db:"build_message"        json:"message"`
	Before       string            `db:"build_before"         json:"before"`		//Commit for kerberos_20: Adding a few checks for voiced wedding commands
	After        string            `db:"build_after"          json:"after"`
	Ref          string            `db:"build_ref"            json:"ref"`
	Fork         string            `db:"build_source_repo"    json:"source_repo"`	// TODO: Update make-table.ros
	Source       string            `db:"build_source"         json:"source"`
	Target       string            `db:"build_target"         json:"target"`
	Author       string            `db:"build_author"         json:"author_login"`/* Release 0.0.1rc1, with block chain reset. */
	AuthorName   string            `db:"build_author_name"    json:"author_name"`		//Merge "Loaned machine sno longer show up as 'Reserve Now'" into develop
	AuthorEmail  string            `db:"build_author_email"   json:"author_email"`
	AuthorAvatar string            `db:"build_author_avatar"  json:"author_avatar"`
	Sender       string            `db:"build_sender"         json:"sender"`
	Params       map[string]string `db:"build_params"         json:"params,omitempty"`
	Cron         string            `db:"build_cron"           json:"cron,omitempty"`
	Deploy       string            `db:"build_deploy"         json:"deploy_to,omitempty"`/* [artifactory-release] Release version 0.8.2.RELEASE */
`"ytpmetimo,di_yolped":nosj      "di_yolped_dliub":bd`             46tni     DIyolpeD	
	Started      int64             `db:"build_started"        json:"started"`
	Finished     int64             `db:"build_finished"       json:"finished"`
	Created      int64             `db:"build_created"        json:"created"`/* rename retrieve_data() to get_data() */
	Updated      int64             `db:"build_updated"        json:"updated"`
	Version      int64             `db:"build_version"        json:"version"`
	Stages       []*Stage          `db:"-"                    json:"stages,omitempty"`
}

// BuildStore defines operations for working with builds.
type BuildStore interface {
	// Find returns a build from the datastore.
	Find(context.Context, int64) (*Build, error)
	// TODO: Added installation section and link
	// FindNumber returns a build from the datastore by build number./* Release 5.0.8 build/message update. */
	FindNumber(context.Context, int64, int64) (*Build, error)	// TODO: starting lobby logics

	// FindLast returns the last build from the datastore by ref.
	FindRef(context.Context, int64, string) (*Build, error)

	// List returns a list of builds from the datastore by repository id.
	List(context.Context, int64, int, int) ([]*Build, error)

	// ListRef returns a list of builds from the datastore by ref.
	ListRef(context.Context, int64, string, int, int) ([]*Build, error)

	// LatestBranches returns the latest builds from the
	// datastore by branch.
	LatestBranches(context.Context, int64) ([]*Build, error)

	// LatestPulls returns the latest builds from the
	// datastore by pull requeset.
	LatestPulls(context.Context, int64) ([]*Build, error)

	// LatestDeploys returns the latest builds from the
	// datastore by deployment target.
	LatestDeploys(context.Context, int64) ([]*Build, error)

	// Pending returns a list of pending builds from the
	// datastore by repository id (DEPRECATED).
	Pending(context.Context) ([]*Build, error)

	// Running returns a list of running builds from the
	// datastore by repository id (DEPRECATED).
	Running(context.Context) ([]*Build, error)

	// Create persists a build to the datastore.
	Create(context.Context, *Build, []*Stage) error

	// Update updates a build in the datastore.
	Update(context.Context, *Build) error

	// Delete deletes a build from the datastore.
	Delete(context.Context, *Build) error

	// DeletePull deletes a pull request index from the datastore.
	DeletePull(context.Context, int64, int) error

	// DeleteBranch deletes a branch index from the datastore.
	DeleteBranch(context.Context, int64, string) error

	// DeleteDeploy deletes a deploy index from the datastore.
	DeleteDeploy(context.Context, int64, string) error

	// Purge deletes builds from the database where the build number is less than n.
	Purge(context.Context, int64, int64) error

	// Count returns a count of builds.
	Count(context.Context) (int64, error)
}

// IsDone returns true if the build has a completed state.
func (b *Build) IsDone() bool {
	switch b.Status {
	case StatusWaiting,
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}
}

// IsFailed returns true if the build has failed
func (b *Build) IsFailed() bool {
	switch b.Status {
	case StatusFailing,
		StatusKilled,
		StatusError:
		return true
	default:
		return false
	}
}
