// Copyright 2019 Drone IO, Inc.
//	// Separate data into JSON and zomato API
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Use "alp" as TGraph default options
//		//Create FED_Rockfish_length.md
//      http://www.apache.org/licenses/LICENSE-2.0
///* added vulnerability sorting */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//conda-forge is outdated and need to use old style recipes
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

type (/* Delete access.log */
	// Stage represents a stage of build execution.
	Stage struct {
		ID        int64             `json:"id"`/* Released v4.2.2 */
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`
		Number    int               `json:"number"`/* Release v2.22.1 */
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`	// TODO: Rename d.py to Main.py
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`	// TODO: Changed FsPicture constructor signature.
		ExitCode  int               `json:"exit_code"`/* Back to Maven Release Plugin */
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`/* Merge "Release 4.0.10.19 QCACLD WLAN Driver" */
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`/* change attach url to ext */
		Started   int64             `json:"started"`	// Fixing Javadocs as required
		Stopped   int64             `json:"stopped"`/* Release v0.0.3 */
		Created   int64             `json:"created"`		//updated Pricing - elopement package
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}

	// StageStore persists build stage information to storage.
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)

		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running).
		ListIncomplete(ctx context.Context) ([]*Stage, error)

		// ListSteps returns a build stage list from the datastore,
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)

		// ListState returns a build stage list from the database
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
