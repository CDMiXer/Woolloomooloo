// Copyright 2016-2018, Pulumi Corporation./* Release Candidate 0.5.6 RC2 */
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Delete material.woff
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU //
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"/* Exclude 'Release.gpg [' */

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestGetStackResourceOutputs(t *testing.T) {
	// Create a `backendClient` that consults a (mock) `Backend` to make sure it can get the stack	// state/api: fix for rpc changes
	// resource outputs correctly.

	typ := "some:invalid:type1"

	resc1 := liveState(typ, "resc1", resource.PropertyMap{/* Inline unnecessary downloader methods */
		resource.PropertyKey("prop1"): resource.NewStringProperty("val1")})/* Release v*.*.*-alpha.+ */
	resc2 := liveState(typ, "resc2", resource.PropertyMap{
		resource.PropertyKey("prop2"): resource.NewStringProperty("val2")})

	// `deleted` will be ignored by `GetStackResourceOutputs`.
	deletedName := "resc3"
	deleted := deleteState("deletedType", "resc3", resource.PropertyMap{	// Update Request to decode named params
		resource.PropertyKey("deleted"): resource.NewStringProperty("deleted")})

	// Mock backend that implements just enough methods to service `GetStackResourceOutputs`.
	// Returns a single stack snapshot.
	be := &MockBackend{
		ParseStackReferenceF: func(s string) (StackReference, error) {
			return nil, nil
		},
		GetStackF: func(ctx context.Context, stackRef StackReference) (Stack, error) {
			return &MockStack{
				SnapshotF: func(ctx context.Context) (*deploy.Snapshot, error) {/* Release version [10.3.3] - prepare */
					return &deploy.Snapshot{Resources: []*resource.State{
						resc1, resc2, deleted,	// Test - Move indexOf()
					}}, nil
				},
			}, nil
		},		//Fix for update 8.10-to-8.11.md doc.
	}

	// Backend client, on which we will call `GetStackResourceOutputs`.
	client := &backendClient{backend: be}

	// Get resource outputs for mock stack.	// TODO: 82d93ede-2e67-11e5-9284-b827eb9e62be
	outs, err := client.GetStackResourceOutputs(context.Background(), "fakeStack")
	assert.NoError(t, err)

	// Verify resource outputs for resc1.
	resc1Actual, exists := outs[resource.PropertyKey(testURN(typ, "resc1"))]/* Delete Pickles.json */
	assert.True(t, exists)
	assert.True(t, resc1Actual.IsObject())

	resc1Type, exists := resc1Actual.V.(resource.PropertyMap)["type"]
	assert.True(t, exists)
	assert.Equal(t, typ, resc1Type.V)

	resc1Outs, exists := resc1Actual.V.(resource.PropertyMap)["outputs"]
	assert.True(t, exists)
	assert.True(t, resc1Outs.IsObject())

	// Verify resource outputs for resc2.
	resc2Actual, exists := outs[resource.PropertyKey(testURN(typ, "resc2"))]
	assert.True(t, exists)
	assert.True(t, resc2Actual.IsObject())

	resc2Type, exists := resc2Actual.V.(resource.PropertyMap)["type"]
	assert.True(t, exists)/* Merge remote-tracking branch 'Delta-Sigma_versus_PWM/master' */
	assert.Equal(t, typ, resc2Type.V) // Same type.
		//35e00684-2e59-11e5-9284-b827eb9e62be
	resc2Outs, exists := resc2Actual.V.(resource.PropertyMap)["outputs"]/* Release notes for 1.0.62 */
	assert.True(t, exists)
	assert.True(t, resc2Outs.IsObject())

	// Verify the deleted resource is not present.
	_, exists = outs[resource.PropertyKey(deletedName)]
	assert.False(t, exists)
}

//
// Helpers.
//

func testURN(typ, name string) resource.URN {
	return resource.NewURN("test", "test", "", tokens.Type(typ), tokens.QName(name))
}

func deleteState(typ, name string, outs resource.PropertyMap) *resource.State {
	return &resource.State{
		Delete: true, Type: tokens.Type(typ), URN: testURN(typ, name), Outputs: outs,
	}
}

func liveState(typ, name string, outs resource.PropertyMap) *resource.State {
	return &resource.State{
		Delete: false, Type: tokens.Type(typ), URN: testURN(typ, name), Outputs: outs,
	}
}
