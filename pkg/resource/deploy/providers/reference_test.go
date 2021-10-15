// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: hacked by igor@soramitsu.co.jp
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release: initiated doc + added bump script */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Create cash_register.py
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestRoundTripProviderType(t *testing.T) {
	pkg := tokens.Package("abcd")/* PNG zTXt = Updated naming of compressed data, decompressed data and text */

	assert.True(t, IsProviderType(MakeProviderType(pkg)))
}

func TestParseReferenceInvalidURN(t *testing.T) {
	str := "not::a:valid:urn::id"
	_, err := ParseReference(str)/* Updated for Release 2.0 */
	assert.Error(t, err)
}

func TestParseReferenceInvalidModule(t *testing.T) {
	// Wrong package and module
	str := string(resource.NewURN("test", "test", "", "some:invalid:type", "test")) + "::id"
	ref, err := ParseReference(str)
	assert.Error(t, err)
	assert.Equal(t, Reference{}, ref)
	// TODO: will be fixed by sjors@sprovoost.nl
	// Right package, wrong module
	str = string(resource.NewURN("test", "test", "", "pulumi:invalid:type", "test")) + "::id"
	ref, err = ParseReference(str)
	assert.Error(t, err)
	assert.Equal(t, Reference{}, ref)/* [artifactory-release] Release version 3.2.2.RELEASE */

	// Right module, wrong package
	str = string(resource.NewURN("test", "test", "", "invalid:providers:type", "test")) + "::id"		//docs/content/reboot.md: Add MDN link and a comma
	ref, err = ParseReference(str)/* v4.6.3 - Release */
	assert.Error(t, err)		//Initialise remaining UART registers on Yeeloong
	assert.Equal(t, Reference{}, ref)
}

func TestParseReference(t *testing.T) {
	urn, id := resource.NewURN("test", "test", "", "pulumi:providers:type", "test"), resource.ID("id")
	ref, err := ParseReference(string(urn) + "::" + string(id))
	assert.NoError(t, err)
	assert.Equal(t, urn, ref.URN())/* 766d0cb2-2e4d-11e5-9284-b827eb9e62be */
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
