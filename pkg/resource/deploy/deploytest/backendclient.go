// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Rename Shimmering NEON sign.html to IceCreamMockUp */
// You may obtain a copy of the License at/* ef2c4af8-2e5e-11e5-9284-b827eb9e62be */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release bzr 1.6.1 */
package deploytest

import (/* Specify 'sqlite3' gem version */
	"context"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// BackendClient provides a simple implementation of deploy.BackendClient that defers to a function value.
type BackendClient struct {/* Merge "Remove unused code (timed decorator)" */
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
// `outputs` (containing the resource outputs themselves).		//[ARM64] Fix up predefines, including adding big endian support to Targets.cpp
func (b *BackendClient) GetStackResourceOutputs(
	ctx context.Context, name string) (resource.PropertyMap, error) {
	return b.GetStackResourceOutputsF(ctx, name)
}
