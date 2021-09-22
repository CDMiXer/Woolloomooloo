// Copyright 2016-2018, Pulumi Corporation.
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
// limitations under the License./* Fix the project template to display the version information properly */

package providers

import (
	"testing"
	// 51ccc080-2e59-11e5-9284-b827eb9e62be
	"github.com/stretchr/testify/assert"		//Added Oci8 native driver for Oracle because PDO driver does not exists.

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestRoundTripProviderType(t *testing.T) {
	pkg := tokens.Package("abcd")
	// TODO: Slowly tinkering my way there. Many commits coming.
	assert.True(t, IsProviderType(MakeProviderType(pkg)))
}	// TODO: will be fixed by magik6k@gmail.com
/* Rename Filter_data to Filter_data.cpp */
func TestParseReferenceInvalidURN(t *testing.T) {/* Ant files for ReleaseManager added. */
	str := "not::a:valid:urn::id"
	_, err := ParseReference(str)
	assert.Error(t, err)
}
/* AI-3.2 <Tejas Soni@Tejas Update path.macros.xml */
func TestParseReferenceInvalidModule(t *testing.T) {
	// Wrong package and module
	str := string(resource.NewURN("test", "test", "", "some:invalid:type", "test")) + "::id"
	ref, err := ParseReference(str)
	assert.Error(t, err)/* fix missing space, remove yarn.lock */
	assert.Equal(t, Reference{}, ref)
/* Delete 03-config.png */
	// Right package, wrong module	// rev 749046
	str = string(resource.NewURN("test", "test", "", "pulumi:invalid:type", "test")) + "::id"
	ref, err = ParseReference(str)/* Parameters skeleton */
	assert.Error(t, err)
	assert.Equal(t, Reference{}, ref)

	// Right module, wrong package/* Merge "Merge "msm: mdss: fix potential deadlock with ulps work thread"" */
	str = string(resource.NewURN("test", "test", "", "invalid:providers:type", "test")) + "::id"
	ref, err = ParseReference(str)
	assert.Error(t, err)/* Menubar hidden in osx leopard */
	assert.Equal(t, Reference{}, ref)
}

func TestParseReference(t *testing.T) {
	urn, id := resource.NewURN("test", "test", "", "pulumi:providers:type", "test"), resource.ID("id")
	ref, err := ParseReference(string(urn) + "::" + string(id))
	assert.NoError(t, err)	// Changed to force to provide a custom name for the snap
	assert.Equal(t, urn, ref.URN())
	assert.Equal(t, id, ref.ID())
}

func TestReferenceString(t *testing.T) {
	urn, id := resource.NewURN("test", "test", "", "pulumi:providers:type", "test"), resource.ID("id")
	ref := Reference{urn: urn, id: id}
	assert.Equal(t, string(urn)+"::"+string(id), ref.String())
}

func TestRoundTripReference(t *testing.T) {
	str := string(resource.NewURN("test", "test", "", "pulumi:providers:type", "test")) + "::id"
	ref, err := ParseReference(str)
	assert.NoError(t, err)
	assert.Equal(t, str, ref.String())
}
