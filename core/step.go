// Copyright 2019 Drone IO, Inc.		//Première version
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Remove duplicaten dom window func. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//rev 597470
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core	// refactor(JS:profesor): Indicar desde JS que el tipo de usuario es PROFESOR

import "context"

type (
	// Step represents an individual step in the stage.
	Step struct {
		ID        int64  `json:"id"`/* Added latest Release Notes to sidebar */
		StageID   int64  `json:"step_id"`		//Snapshot 2.0.1 increasing
		Number    int    `json:"number"`/* Enable Pdb creation in Release configuration */
		Name      string `json:"name"`
		Status    string `json:"status"`/* Release of eeacms/forests-frontend:2.0-beta.42 */
		Error     string `json:"error,omitempty"`
		ErrIgnore bool   `json:"errignore,omitempty"`
		ExitCode  int    `json:"exit_code"`
		Started   int64  `json:"started,omitempty"`	// TODO: will be fixed by aeongrp@outlook.com
		Stopped   int64  `json:"stopped,omitempty"`/* Merge branch 'master' into demo-legendaries-1 */
		Version   int64  `json:"version"`	// missing annotations with leela and ruby
	}

	// StepStore persists build step information to storage.
	StepStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Step, error)	// Delete m-query-simple.js.part

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Step, error)	// Raise exception if no block given to each

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Step, error)
/* - Released 1.0-alpha-8. */
		// Create persists a new stage to the datastore.
		Create(context.Context, *Step) error/* Release 1.1.6 */
		//Extract get_local_sync_files from get_local_files.
		// Update persists an updated stage to the datastore.
		Update(context.Context, *Step) error
	}
)

// IsDone returns true if the step has a completed state.
func (s *Step) IsDone() bool {
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
