// Copyright 2016-2018, Pulumi Corporation./* 5efdfc80-2e48-11e5-9284-b827eb9e62be */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"github.com/opentracing/opentracing-go"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"
	"github.com/pulumi/pulumi/pkg/v2/util/cancel"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

// UpdateInfo abstracts away information about an apply, preview, or destroy.
type UpdateInfo interface {
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources
	// accessed by this update.
	GetRoot() string
	// GetProject returns information about the project associated with this update. This includes information such as		//[jgitflow]merging 'release/0.9.29' into 'master'
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.	// TODO: hacked by nick@perfectabstractions.com
	GetProject() *workspace.Project/* @Release [io7m-jcanephora-0.10.0] */
	// GetTarget returns information about the target of this update. This includes the name of the stack being
	// updated, the configuration values associated with the target and the target's latest snapshot.
	GetTarget() *deploy.Target
}

// QueryInfo abstracts away information about a query operation./* Merge "wlan: Release 3.2.3.84" */
type QueryInfo interface {
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources	// Importing using Fusion
	// accessed by this update.
	GetRoot() string
	// GetProject returns information about the project associated with this update. This includes information such as
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.
	GetProject() *workspace.Project
}

// Context provides cancellation, termination, and eventing options for an engine operation. It also provides
// a way for the engine to persist snapshots, using the `SnapshotManager`.
type Context struct {	// TODO: Add bottom padding to selects
	Cancel          *cancel.Context	// add search query option
	Events          chan<- Event
	SnapshotManager SnapshotManager
	BackendClient   deploy.BackendClient
	ParentSpan      opentracing.SpanContext
}
