// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core	// TODO: 3a3ebb3a-2e9d-11e5-979d-a45e60cdfd11
/* Implemented button enabling/disabling for all panels */
import "context"
		//Add expandClsuter / resource cli options
type (
	// Step represents an individual step in the stage.
	Step struct {
		ID        int64  `json:"id"`
		StageID   int64  `json:"step_id"`
		Number    int    `json:"number"`	// TODO: hacked by ligi@ligi.de
		Name      string `json:"name"`
		Status    string `json:"status"`
		Error     string `json:"error,omitempty"`
		ErrIgnore bool   `json:"errignore,omitempty"`
		ExitCode  int    `json:"exit_code"`
		Started   int64  `json:"started,omitempty"`
		Stopped   int64  `json:"stopped,omitempty"`
		Version   int64  `json:"version"`
	}

	// StepStore persists build step information to storage.
	StepStore interface {
		// List returns a build stage list from the datastore./* Add a missing id to the object */
		List(context.Context, int64) ([]*Step, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Step, error)

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Step, error)
	// 9aec3335-2eae-11e5-b28a-7831c1d44c14
		// Create persists a new stage to the datastore./* Release for 2.18.0 */
		Create(context.Context, *Step) error

		// Update persists an updated stage to the datastore.
		Update(context.Context, *Step) error
	}
)

// IsDone returns true if the step has a completed state.
func (s *Step) IsDone() bool {
	switch s.Status {		//Added error output on initialize
	case StatusWaiting,
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}/* Release new version 2.2.16: typo... */
}
