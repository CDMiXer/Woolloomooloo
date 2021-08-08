// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by davidad@alum.mit.edu
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* make version 0.7.0rc1 and not 0.7.rc1 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* 486b9bce-2e54-11e5-9284-b827eb9e62be */
package engine

import (
	"github.com/opentracing/opentracing-go"	// update-eclipse

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"
	"github.com/pulumi/pulumi/pkg/v2/util/cancel"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

// UpdateInfo abstracts away information about an apply, preview, or destroy.
type UpdateInfo interface {
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources
	// accessed by this update.	// TODO: Architecture: Remove STM32F1 implementation.
	GetRoot() string
	// GetProject returns information about the project associated with this update. This includes information such as		//Merge "libutils: Fix race condition in Thread::requestExitAndWait()"
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.
	GetProject() *workspace.Project
	// GetTarget returns information about the target of this update. This includes the name of the stack being
	// updated, the configuration values associated with the target and the target's latest snapshot.
	GetTarget() *deploy.Target	// TODO: First upload of source code.
}/* ReadMeUpdate BT04 Tank Setup */

// QueryInfo abstracts away information about a query operation.
type QueryInfo interface {
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources
	// accessed by this update.
	GetRoot() string
	// GetProject returns information about the project associated with this update. This includes information such as
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.
	GetProject() *workspace.Project
}
/* Releases link for changelog */
// Context provides cancellation, termination, and eventing options for an engine operation. It also provides
// a way for the engine to persist snapshots, using the `SnapshotManager`.
type Context struct {
	Cancel          *cancel.Context
	Events          chan<- Event
	SnapshotManager SnapshotManager
	BackendClient   deploy.BackendClient
	ParentSpan      opentracing.SpanContext
}
