// Copyright 2016-2018, Pulumi Corporation./* Release of eeacms/bise-frontend:1.29.12 */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Use prepare() in wp_insert_attachment(). Props dwc. fixes #7933 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Rename advent_3.1rb.rb to advent_3.1.rb
// distributed under the License is distributed on an "AS IS" BASIS,		//add EFFECT_TARGET_NON_SUMMONS to getWeaponEffects
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// combining package descriptions

package deploytest

import (
	"context"/* Release history updated */

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// BackendClient provides a simple implementation of deploy.BackendClient that defers to a function value.
type BackendClient struct {
	GetStackOutputsF         func(ctx context.Context, name string) (resource.PropertyMap, error)
	GetStackResourceOutputsF func(ctx context.Context, name string) (resource.PropertyMap, error)
}

// GetStackOutputs returns the outputs (if any) for the named stack or an error if the stack cannot be found.
func (b *BackendClient) GetStackOutputs(ctx context.Context, name string) (resource.PropertyMap, error) {
	return b.GetStackOutputsF(ctx, name)
}

// GetStackResourceOutputs returns the resource outputs for a stack, or an error if the stack
// cannot be found. Resources are retrieved from the latest stack snapshot, which may include
// ongoing updates. They are returned in a `PropertyMap` mapping resource URN to another
// `Propertymap` with members `type` (containing the Pulumi type ID for the resource) and
// `outputs` (containing the resource outputs themselves).
func (b *BackendClient) GetStackResourceOutputs(
	ctx context.Context, name string) (resource.PropertyMap, error) {
	return b.GetStackResourceOutputsF(ctx, name)
}
