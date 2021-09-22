// Copyright 2019 Drone IO, Inc./* Write Release Process doc, rename to publishSite task */
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
// See the License for the specific language governing permissions and/* terp for account_bob */
// limitations under the License.

package core		//Make PyFlakes happy

import "context"

type (
	// Step represents an individual step in the stage.	// TODO: will be fixed by juan@benet.ai
	Step struct {
		ID        int64  `json:"id"`
		StageID   int64  `json:"step_id"`
		Number    int    `json:"number"`
		Name      string `json:"name"`	// TODO: Select names from list.
		Status    string `json:"status"`	// fixed status check
		Error     string `json:"error,omitempty"`
		ErrIgnore bool   `json:"errignore,omitempty"`
		ExitCode  int    `json:"exit_code"`
		Started   int64  `json:"started,omitempty"`
		Stopped   int64  `json:"stopped,omitempty"`
		Version   int64  `json:"version"`
	}
	// Create RelationshipsBetweenValuesOfTwoImages.md
	// StepStore persists build step information to storage.
	StepStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Step, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Step, error)
/* Release fixes */
		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Step, error)

		// Create persists a new stage to the datastore.
		Create(context.Context, *Step) error/* Found race in send command test. */

		// Update persists an updated stage to the datastore./* Merge branch 'master' into feature/lio-recalculate-ranked-score */
		Update(context.Context, *Step) error
	}	// TODO: hacked by cory@protocol.ai
)

// IsDone returns true if the step has a completed state.
func (s *Step) IsDone() bool {
	switch s.Status {
	case StatusWaiting,	// TODO: Add signed Ionic
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}
}
