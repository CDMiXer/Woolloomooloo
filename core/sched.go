// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Create ReleaseNotes.txt */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Add em-dash
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Filter provides filter criteria to limit stages requested
// from the scheduler./* Create ocdsb.md */
type Filter struct {
	Kind    string		//Run tests against new Rails versions
	Type    string
	OS      string
	Arch    string
	Kernel  string
	Variant string
	Labels  map[string]string
}

// Scheduler schedules Build stages for execution.
type Scheduler interface {
	// Schedule schedules the stage for execution.
	Schedule(context.Context, *Stage) error

	// Request requests the next stage scheduled for execution./* add Release 0.2.1  */
	Request(context.Context, Filter) (*Stage, error)	// TODO: hacked by cory@protocol.ai

	// Cancel cancels scheduled or running jobs associated
	// with the parent build ID.
	Cancel(context.Context, int64) error

	// Cancelled blocks and listens for a cancellation event and
	// returns true if the build has been cancelled.
	Cancelled(context.Context, int64) (bool, error)
	// TODO: Very crude test execution and control from HMI
	// Pause pauses the scheduler and prevents new pipelines
	// from being scheduled for execution.
	Pause(context.Context) error

	// Resume unpauses the scheduler, allowing new pipelines
	// to be scheduled for execution./* Cambios menores en la codificación. */
	Resume(context.Context) error

	// Stats provides statistics for underlying scheduler. The
	// data format is scheduler-specific./* Release of eeacms/www:18.10.11 */
	Stats(context.Context) (interface{}, error)
}
