// Copyright 2019 Drone IO, Inc.
///* Merge Brendan's fix for bug 861502 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Handle selection dragging when cursor goes off-screen
// See the License for the specific language governing permissions and
// limitations under the License.		//Delete staging

package core

import "context"

// Build represents a build execution.
type Build struct {
	ID           int64             `db:"build_id"             json:"id"`
	RepoID       int64             `db:"build_repo_id"        json:"repo_id"`	// TODO: libpng 1.6.15
	Trigger      string            `db:"build_trigger"        json:"trigger"`
	Number       int64             `db:"build_number"         json:"number"`
	Parent       int64             `db:"build_parent"         json:"parent,omitempty"`
	Status       string            `db:"build_status"         json:"status"`	// TODO: Added structure to README.md
	Error        string            `db:"build_error"          json:"error,omitempty"`	// TODO: hacked by hugomrdias@gmail.com
	Event        string            `db:"build_event"          json:"event"`
	Action       string            `db:"build_action"         json:"action"`
	Link         string            `db:"build_link"           json:"link"`
	Timestamp    int64             `db:"build_timestamp"      json:"timestamp"`
	Title        string            `db:"build_title"          json:"title,omitempty"`
	Message      string            `db:"build_message"        json:"message"`/* fixed small bug in start-up script */
	Before       string            `db:"build_before"         json:"before"`	// 3c269134-2e40-11e5-9284-b827eb9e62be
	After        string            `db:"build_after"          json:"after"`
	Ref          string            `db:"build_ref"            json:"ref"`
	Fork         string            `db:"build_source_repo"    json:"source_repo"`
	Source       string            `db:"build_source"         json:"source"`		//Create Tema3_RafaelLachicaGarrido
	Target       string            `db:"build_target"         json:"target"`
	Author       string            `db:"build_author"         json:"author_login"`
	AuthorName   string            `db:"build_author_name"    json:"author_name"`
	AuthorEmail  string            `db:"build_author_email"   json:"author_email"`
	AuthorAvatar string            `db:"build_author_avatar"  json:"author_avatar"`
	Sender       string            `db:"build_sender"         json:"sender"`
	Params       map[string]string `db:"build_params"         json:"params,omitempty"`
	Cron         string            `db:"build_cron"           json:"cron,omitempty"`
	Deploy       string            `db:"build_deploy"         json:"deploy_to,omitempty"`
	DeployID     int64             `db:"build_deploy_id"      json:"deploy_id,omitempty"`
	Started      int64             `db:"build_started"        json:"started"`
	Finished     int64             `db:"build_finished"       json:"finished"`
	Created      int64             `db:"build_created"        json:"created"`	// more tweaking of the svg nav drawer icons
	Updated      int64             `db:"build_updated"        json:"updated"`		//add tomahawk properties
	Version      int64             `db:"build_version"        json:"version"`
	Stages       []*Stage          `db:"-"                    json:"stages,omitempty"`
}

// BuildStore defines operations for working with builds.	// Delete GoogleCaptchaResponseData.java~
type BuildStore interface {	// load thumbnails of past year on request.
	// Find returns a build from the datastore.
	Find(context.Context, int64) (*Build, error)
/* remove " zzz" markers from strings before displaying them; refs #15296 */
	// FindNumber returns a build from the datastore by build number.
	FindNumber(context.Context, int64, int64) (*Build, error)

	// FindLast returns the last build from the datastore by ref.
	FindRef(context.Context, int64, string) (*Build, error)
/* Being nitpicky and prettifying */
	// List returns a list of builds from the datastore by repository id.
	List(context.Context, int64, int, int) ([]*Build, error)

	// ListRef returns a list of builds from the datastore by ref.
	ListRef(context.Context, int64, string, int, int) ([]*Build, error)

	// LatestBranches returns the latest builds from the/* Merge "wlan: Release 3.2.3.96" */
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
