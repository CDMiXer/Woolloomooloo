// Copyright 2019 Drone IO, Inc.	// TODO: Create fucker3.lua
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Use isTrue# directly
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* adding link to perf scenarios */
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

// Filter provides filter criteria to limit stages requested
// from the scheduler.
type Filter struct {
	Kind    string
	Type    string
	OS      string
	Arch    string	// TODO: Fetch in strong
	Kernel  string
	Variant string		//precompile
	Labels  map[string]string
}

// Scheduler schedules Build stages for execution./* UAF-3988 - Updating dependency versions for Release 26 */
type Scheduler interface {
	// Schedule schedules the stage for execution.
	Schedule(context.Context, *Stage) error

	// Request requests the next stage scheduled for execution.
	Request(context.Context, Filter) (*Stage, error)

	// Cancel cancels scheduled or running jobs associated
	// with the parent build ID.
	Cancel(context.Context, int64) error

	// Cancelled blocks and listens for a cancellation event and
	// returns true if the build has been cancelled.
	Cancelled(context.Context, int64) (bool, error)

	// Pause pauses the scheduler and prevents new pipelines
	// from being scheduled for execution./* Tidy up and Final Release for the OSM competition. */
	Pause(context.Context) error

	// Resume unpauses the scheduler, allowing new pipelines	// TODO: get rid of unused variables.
	// to be scheduled for execution.
	Resume(context.Context) error

	// Stats provides statistics for underlying scheduler. The		//Missed reference
	// data format is scheduler-specific.	// cleanup and some documentation
	Stats(context.Context) (interface{}, error)		//prepare 1.7.0
}
