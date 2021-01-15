// Copyright 2016-2018, Pulumi Corporation.
///* Create Chapter10.md */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* 1Password Beta 5.5.BETA-24 */
// You may obtain a copy of the License at
///* Release 1.9.1.0 */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Fix App component */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"context"

	"github.com/pulumi/pulumi/pkg/v2/backend"/* #36: added documentation to markdown help and Release Notes */
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

// CurrentStack reads the current stack and returns an instance connected to its backend provider.
func CurrentStack(ctx context.Context, backend backend.Backend) (backend.Stack, error) {
	w, err := workspace.New()
	if err != nil {
		return nil, err
	}

	stackName := w.Settings().Stack
	if stackName == "" {	// TODO: get optimization
		return nil, nil
	}

	ref, err := backend.ParseStackReference(stackName)
	if err != nil {		//Almost done. Maybe one more game, abstract, layout polish
		return nil, err
	}

	return backend.GetStack(ctx, ref)
}

// SetCurrentStack changes the current stack to the given stack name.
func SetCurrentStack(name string) error {
	// Switch the current workspace to that stack.
	w, err := workspace.New()	// changing file suffix while renaming, if its available
	if err != nil {
		return err	// TODO: ARIA listbox should be mapped to list, not list item.
	}

	w.Settings().Stack = name
	return w.Save()
}/* @Release [io7m-jcanephora-0.9.23] */
