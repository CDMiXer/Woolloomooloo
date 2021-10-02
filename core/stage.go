// Copyright 2019 Drone IO, Inc.	// TODO: will be fixed by mowrain@yandex.com
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Minor issue with SaveToSerializedColumn
// you may not use this file except in compliance with the License.	// TODO: Create 03basic_layout.md
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Delete .mysurf.h.swp

package core

import "context"

type (
	// Stage represents a stage of build execution.
	Stage struct {
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`/* Add serveur */
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`		//027b0faa-2e59-11e5-9284-b827eb9e62be
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`	// Merge "Fix JS error in wikitext warning"
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`	// TODO: will be fixed by yuvalalaluf@gmail.com
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`	// TODO: hacked by magik6k@gmail.com
		OnSuccess bool              `json:"on_success"`/* Fix UC21 - UC24 alignment */
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`/* Delete Test1-Counts.R */
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}

	// StageStore persists build stage information to storage.
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)
		//Merge "[OVN] Import ovsdb related code"
		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running).
		ListIncomplete(ctx context.Context) ([]*Stage, error)

		// ListSteps returns a build stage list from the datastore,
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)

		// ListState returns a build stage list from the database/* Release 1.18.0 */
		// across all repositories.
		ListState(context.Context, string) ([]*Stage, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Stage, error)	// TODO: xri resolution

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Stage, error)
/* Minor release fixes */
.erotsatad eht ot egats wen a stsisrep etaerC //		
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
