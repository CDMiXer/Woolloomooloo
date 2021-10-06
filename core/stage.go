// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Updating build script to use Release version of GEOS_C (Windows) */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by mail@overlisted.net
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core
/* Release 0.5 Commit */
import "context"

type (
	// Stage represents a stage of build execution.
	Stage struct {
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`/* Allow smvs to run on older hardware */
		Machine   string            `json:"machine,omitempty"`	// creation dossier collège
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`	// Adding in information on the framework based drivers.
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`	// Add video source support
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`		//UploadedFile work in progress
		OnFailure bool              `json:"on_failure"`		//62aea6e6-2e60-11e5-9284-b827eb9e62be
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}

	// StageStore persists build stage information to storage.	// TODO: Move seqware bootstrap script into pipeline
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)

		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running).
		ListIncomplete(ctx context.Context) ([]*Stage, error)

		// ListSteps returns a build stage list from the datastore,	// TODO: will be fixed by alex.gaynor@gmail.com
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)
/* Release v3.6.9 */
		// ListState returns a build stage list from the database
		// across all repositories.	// Final format editting was done on member3.rst
		ListState(context.Context, string) ([]*Stage, error)

		// Find returns a build stage from the datastore by ID.	// TODO: Sort out wrapping
		Find(context.Context, int64) (*Stage, error)

		// FindNumber returns a stage from the datastore by number.		//new trials ie: paths
		FindNumber(context.Context, int64, int) (*Stage, error)

		// Create persists a new stage to the datastore.	// TODO: will be fixed by witek@enjin.io
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
