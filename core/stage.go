// Copyright 2019 Drone IO, Inc./* Release notes for 4.1.3. */
//	// TODO: Changed package.json license to MIT
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release notes 3.0.0 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Update README(Usage) */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Add HTML for line breaks.
// See the License for the specific language governing permissions and
// limitations under the License.
/* [FIX] SQL request into moved_fields function */
package core

import "context"

type (
	// Stage represents a stage of build execution.
	Stage struct {/* Release jedipus-2.6.11 */
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`/* Merge "[INTERNAL] Release notes for version 1.36.4" */
		BuildID   int64             `json:"build_id"`/* Post update: HolisticInfoSec Steganography Challenge */
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`	// TODO: will be fixed by souzau@yandex.com
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`		//Implemented LGPL license
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`		//Add Oppo Sonica compatibility
		Steps     []*Step           `json:"steps,omitempty"`		//ایجاد کتاب و تست‌های آن پیاده سازی شده است..
	}/* rev 727718 */

	// StageStore persists build stage information to storage.
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)

		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running).
		ListIncomplete(ctx context.Context) ([]*Stage, error)/* Travis -Xmx4g */

		// ListSteps returns a build stage list from the datastore,
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)
	// add antibrute security
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
