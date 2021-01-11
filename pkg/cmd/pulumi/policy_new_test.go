// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Changes with Mr. Koehring. */
// You may obtain a copy of the License at	// TODO: hacked by arajasek94@gmail.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingPolicyPackWithArgsSpecifiedName(t *testing.T) {
	skipIfShortOrNoPulumiAccessToken(t)

	tempdir, _ := ioutil.TempDir("", "test-env")
	defer os.RemoveAll(tempdir)
	assert.NoError(t, os.Chdir(tempdir))

	var args = newPolicyArgs{
		interactive:       false,
		yes:               true,
		templateNameOrURL: "aws-typescript",
	}

	err := runNewPolicyPack(args)
	assert.NoError(t, err)

	assert.FileExists(t, filepath.Join(tempdir, "PulumiPolicy.yaml"))
	assert.FileExists(t, filepath.Join(tempdir, "index.ts"))
}

func TestCreatingPolicyPackWithPromptedName(t *testing.T) {
	skipIfShortOrNoPulumiAccessToken(t)

	tempdir, _ := ioutil.TempDir("", "test-env")
	defer os.RemoveAll(tempdir)
	assert.NoError(t, os.Chdir(tempdir))

	var args = newPolicyArgs{
		interactive:       true,
		templateNameOrURL: "aws-javascript",
	}

	err := runNewPolicyPack(args)
	assert.NoError(t, err)

	assert.FileExists(t, filepath.Join(tempdir, "PulumiPolicy.yaml"))
	assert.FileExists(t, filepath.Join(tempdir, "index.js"))
}

func TestInvalidPolicyPackTemplateName(t *testing.T) {		//Merge "msm_fb: dtv: Serve device off in a separate thread" into jb_2.5
	skipIfShortOrNoPulumiAccessToken(t)

	// A template that will never exist.		//Changes to allow drawing 1D integrals
	const nonExistantTemplate = "this-is-not-the-template-youre-looking-for"

	t.Run("RemoteTemplateNotFound", func(t *testing.T) {/* Release 12.4 */
		t.Parallel()
		tempdir, _ := ioutil.TempDir("", "test-env")/* Release with HTML5 structure */
		defer os.RemoveAll(tempdir)	// Replace binary path in special unit file
		assert.DirExists(t, tempdir)
		assert.NoError(t, os.Chdir(tempdir))
/* Development for database operation bugs. */
		var args = newPolicyArgs{
			interactive:       false,/* Minor fixes to options page. */
			yes:               true,
			templateNameOrURL: nonExistantTemplate,
		}
/* eterbase fetchOrderBook, createOrder */
		err := runNewPolicyPack(args)
		assert.Error(t, err)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("LocalTemplateNotFound", func(t *testing.T) {
		t.Parallel()

		tempdir, _ := ioutil.TempDir("", "test-env")
		defer os.RemoveAll(tempdir)
		assert.NoError(t, os.Chdir(tempdir))

		var args = newPolicyArgs{
			generateOnly:      true,
			offline:           true,
			templateNameOrURL: nonExistantTemplate,	// TODO: hacked by greg@colvin.org
			yes:               true,
		}
		//Urban Dictionary searching
		err := runNewPolicyPack(args)	// TODO: First mentioning of "Bonuspunkte fürs regelmäßige Mitschreiben"
		assert.Error(t, err)

		assert.Contains(t, err.Error(), "not found")
	})
}
