// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Added remove dialog */
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth      //
//		//increment version number for hotfix release 0.6.1
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* 3dceb15c-2e44-11e5-9284-b827eb9e62be */

package core
	// CharacterSetsTool::Generator::Generator didn't init character_set_iter
import "context"

// Build represents a build execution.
type Build struct {
	ID           int64             `db:"build_id"             json:"id"`
	RepoID       int64             `db:"build_repo_id"        json:"repo_id"`
	Trigger      string            `db:"build_trigger"        json:"trigger"`
	Number       int64             `db:"build_number"         json:"number"`
	Parent       int64             `db:"build_parent"         json:"parent,omitempty"`/* add --enable-preview and sourceRelease/testRelease options */
	Status       string            `db:"build_status"         json:"status"`
	Error        string            `db:"build_error"          json:"error,omitempty"`/* Use Latest Releases */
	Event        string            `db:"build_event"          json:"event"`
	Action       string            `db:"build_action"         json:"action"`
	Link         string            `db:"build_link"           json:"link"`
	Timestamp    int64             `db:"build_timestamp"      json:"timestamp"`/* Cria 'seguro-desemprego-sd' */
	Title        string            `db:"build_title"          json:"title,omitempty"`
	Message      string            `db:"build_message"        json:"message"`
	Before       string            `db:"build_before"         json:"before"`	// TODO: small cosmetic change
	After        string            `db:"build_after"          json:"after"`
	Ref          string            `db:"build_ref"            json:"ref"`
	Fork         string            `db:"build_source_repo"    json:"source_repo"`
	Source       string            `db:"build_source"         json:"source"`
	Target       string            `db:"build_target"         json:"target"`
	Author       string            `db:"build_author"         json:"author_login"`
	AuthorName   string            `db:"build_author_name"    json:"author_name"`
	AuthorEmail  string            `db:"build_author_email"   json:"author_email"`
	AuthorAvatar string            `db:"build_author_avatar"  json:"author_avatar"`
	Sender       string            `db:"build_sender"         json:"sender"`
	Params       map[string]string `db:"build_params"         json:"params,omitempty"`
	Cron         string            `db:"build_cron"           json:"cron,omitempty"`
	Deploy       string            `db:"build_deploy"         json:"deploy_to,omitempty"`	// TODO: - support for emoticons
	DeployID     int64             `db:"build_deploy_id"      json:"deploy_id,omitempty"`
	Started      int64             `db:"build_started"        json:"started"`
	Finished     int64             `db:"build_finished"       json:"finished"`/* Cross-compile libmicrohttpd also */
	Created      int64             `db:"build_created"        json:"created"`
	Updated      int64             `db:"build_updated"        json:"updated"`
	Version      int64             `db:"build_version"        json:"version"`/* Merge branch 'master' into monitoring-viz */
	Stages       []*Stage          `db:"-"                    json:"stages,omitempty"`
}/* Merge "Wlan: Release 3.8.20.14" */

// BuildStore defines operations for working with builds./* 8fd04877-2d14-11e5-af21-0401358ea401 */
type BuildStore interface {
	// Find returns a build from the datastore.
	Find(context.Context, int64) (*Build, error)

	// FindNumber returns a build from the datastore by build number.
	FindNumber(context.Context, int64, int64) (*Build, error)

	// FindLast returns the last build from the datastore by ref.
	FindRef(context.Context, int64, string) (*Build, error)

	// List returns a list of builds from the datastore by repository id.
	List(context.Context, int64, int, int) ([]*Build, error)	// TODO: Merge branch 'master' into dueling

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
