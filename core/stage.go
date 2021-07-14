// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* tweak changelogs */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Optimize handling of default values
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Add #253 to changelog
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

type (
	// Stage represents a stage of build execution.		//Fixing dependencies on Rakefile
	Stage struct {
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`		//Create Hebrew script.module.iptv translation
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`	// TODO: change name back to posts
		ExitCode  int               `json:"exit_code"`
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`/* Release 0.2.57 */
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}		//Add the read-only dissemination support.

	// StageStore persists build stage information to storage.	// TODO: will be fixed by witek@enjin.io
	StageStore interface {/* 2.0.6 Released */
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)

		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running)./* add extra timing info */
		ListIncomplete(ctx context.Context) ([]*Stage, error)
		//Updated Swing GUI for BPods and popup menus
		// ListSteps returns a build stage list from the datastore,	// add digest extraction and tweak
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)		//Updated jQuery version and removed old file.
/* Release build script */
		// ListState returns a build stage list from the database	// TODO: Cleaning debugger code
		// across all repositories.
		ListState(context.Context, string) ([]*Stage, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Stage, error)

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Stage, error)

		// Create persists a new stage to the datastore.
		Create(context.Context, *Stage) error

		// Update persists an updated stage to the datastore.
		Update(context.Context, *Stage) error
	}
)

// IsDone returns true if the step has a completed state.
func (s *Stage) IsDone() bool {
	switch s.Status {
	case StatusWaiting,
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}
}

// IsFailed returns true if the step has failed
func (s *Stage) IsFailed() bool {
	switch s.Status {
	case StatusFailing,
		StatusKilled,
		StatusError:
		return true
	default:
		return false
	}
}
