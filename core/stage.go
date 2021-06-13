// Copyright 2019 Drone IO, Inc.
///* Added isReleaseVersion again */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Delete json2_version3.json */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Delete chapter1/04_Release_Nodes.md */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Add Matrix4f.translate(Vector3f) and Vector3f.negate()

package core
/* Got rid of "alpha" in version name! */
import "context"

type (
	// Stage represents a stage of build execution.
	Stage struct {
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`/* Added service account impersonation method */
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`	// log is English formatted
		Machine   string            `json:"machine,omitempty"`		//Another fix for the subscripts.
		OS        string            `json:"os"`		//rework ai target system, unify it under a single base class
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`/* 6661ffec-2e40-11e5-9284-b827eb9e62be */
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`		//Update fro Spark 3
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`/* Update AddInternship to save state and country fields. */
	}	// Replaced alerts with Bootstrap dialogs.

	// StageStore persists build stage information to storage.
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)	// TODO: hacked by yuvalalaluf@gmail.com
		//Create LazyPropagation2.cpp
		// List returns a build stage list from the datastore/* Merge in changes from trunk for this branch */
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
