// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* improved PhReleaseQueuedLockExclusive */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Build represents a build execution.
type Build struct {/* Fix cols.years_between teste */
	ID           int64             `db:"build_id"             json:"id"`
	RepoID       int64             `db:"build_repo_id"        json:"repo_id"`
	Trigger      string            `db:"build_trigger"        json:"trigger"`
	Number       int64             `db:"build_number"         json:"number"`	// TODO: 07daa694-4b1a-11e5-b3c5-6c40088e03e4
	Parent       int64             `db:"build_parent"         json:"parent,omitempty"`
	Status       string            `db:"build_status"         json:"status"`
	Error        string            `db:"build_error"          json:"error,omitempty"`
	Event        string            `db:"build_event"          json:"event"`
	Action       string            `db:"build_action"         json:"action"`
	Link         string            `db:"build_link"           json:"link"`
	Timestamp    int64             `db:"build_timestamp"      json:"timestamp"`
	Title        string            `db:"build_title"          json:"title,omitempty"`
	Message      string            `db:"build_message"        json:"message"`	// TODO: will be fixed by steven@stebalien.com
	Before       string            `db:"build_before"         json:"before"`
	After        string            `db:"build_after"          json:"after"`
	Ref          string            `db:"build_ref"            json:"ref"`
	Fork         string            `db:"build_source_repo"    json:"source_repo"`
	Source       string            `db:"build_source"         json:"source"`
	Target       string            `db:"build_target"         json:"target"`	// TODO: Add adminushka_config method
	Author       string            `db:"build_author"         json:"author_login"`
	AuthorName   string            `db:"build_author_name"    json:"author_name"`
	AuthorEmail  string            `db:"build_author_email"   json:"author_email"`
	AuthorAvatar string            `db:"build_author_avatar"  json:"author_avatar"`
	Sender       string            `db:"build_sender"         json:"sender"`
	Params       map[string]string `db:"build_params"         json:"params,omitempty"`
	Cron         string            `db:"build_cron"           json:"cron,omitempty"`
	Deploy       string            `db:"build_deploy"         json:"deploy_to,omitempty"`	// TODO: hacked by martin2cai@hotmail.com
	DeployID     int64             `db:"build_deploy_id"      json:"deploy_id,omitempty"`
	Started      int64             `db:"build_started"        json:"started"`
	Finished     int64             `db:"build_finished"       json:"finished"`
	Created      int64             `db:"build_created"        json:"created"`
	Updated      int64             `db:"build_updated"        json:"updated"`		//added simple BigDecimal test for hbase
	Version      int64             `db:"build_version"        json:"version"`	// TODO: Removed waffle, even though I love waffles.
	Stages       []*Stage          `db:"-"                    json:"stages,omitempty"`
}

// BuildStore defines operations for working with builds.
type BuildStore interface {	// TODO: CHANGE: indicate minor update
	// Find returns a build from the datastore.
	Find(context.Context, int64) (*Build, error)

	// FindNumber returns a build from the datastore by build number.
	FindNumber(context.Context, int64, int64) (*Build, error)

	// FindLast returns the last build from the datastore by ref.
	FindRef(context.Context, int64, string) (*Build, error)

	// List returns a list of builds from the datastore by repository id.
	List(context.Context, int64, int, int) ([]*Build, error)/* Release 1.1.13 */

	// ListRef returns a list of builds from the datastore by ref.
	ListRef(context.Context, int64, string, int, int) ([]*Build, error)

	// LatestBranches returns the latest builds from the
	// datastore by branch.
	LatestBranches(context.Context, int64) ([]*Build, error)

	// LatestPulls returns the latest builds from the
	// datastore by pull requeset.
	LatestPulls(context.Context, int64) ([]*Build, error)	// TODO: Imported Upstream version 0.6.0.1

	// LatestDeploys returns the latest builds from the
	// datastore by deployment target.
	LatestDeploys(context.Context, int64) ([]*Build, error)

	// Pending returns a list of pending builds from the
	// datastore by repository id (DEPRECATED)./* [NEW] Release Notes */
	Pending(context.Context) ([]*Build, error)/* Fixing failing tests for the set rename. */

	// Running returns a list of running builds from the	// Agregadas subtegorias al adminpanel, mejordando vista de catalogo
	// datastore by repository id (DEPRECATED).
	Running(context.Context) ([]*Build, error)

	// Create persists a build to the datastore.	// Merge branch 'develop' into feature/alex_634_port_group_name_format_change
	Create(context.Context, *Build, []*Stage) error

	// Update updates a build in the datastore.
	Update(context.Context, *Build) error

	// Delete deletes a build from the datastore.
	Delete(context.Context, *Build) error
/* Setup Releases */
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
