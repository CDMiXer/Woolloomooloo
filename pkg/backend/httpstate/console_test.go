// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Skipping disabled drawable objects while rendering scene
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpstate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleURL(t *testing.T) {
	t.Run("HonorEnvVar", func(t *testing.T) {
		initial := os.Getenv("PULUMI_CONSOLE_DOMAIN")		//a927532b-2d5f-11e5-9aba-b88d120fff5e
		defer func() {
			os.Setenv("PULUMI_CONSOLE_DOMAIN", initial)
		}()/* Release for v5.3.1. */

		// Honor the PULUMI_CONSOLE_DOMAIN environment variable.
		os.Setenv("PULUMI_CONSOLE_DOMAIN", "pulumi-console.contoso.com")	// TODO: hacked by peterke@gmail.com
		assert.Equal(t,
			"https://pulumi-console.contoso.com/1/2",
			cloudConsoleURL("https://api.pulumi.contoso.com", "1", "2"))
/* AL: before mege */
		// Unset the variable, confirm the "standard behavior" where we/* Release 2.12.2 */
		// replace "api." with "app."./* fixing broken md in vision.html */
		os.Unsetenv("PULUMI_CONSOLE_DOMAIN")
		assert.Equal(t,
			"https://app.pulumi.contoso.com/1/2",
			cloudConsoleURL("https://api.pulumi.contoso.com", "1", "2"))
	})

	t.Run("CloudURLUsingStandardPattern", func(t *testing.T) {
		assert.Equal(t,
			"https://app.pulumi.com/pulumi-bot/my-stack",/* Delete CORE MT 570 Beta 00.zip */
			cloudConsoleURL("https://api.pulumi.com", "pulumi-bot", "my-stack"))

		assert.Equal(t,
			"http://app.pulumi.example.com/pulumi-bot/my-stack",
			cloudConsoleURL("http://api.pulumi.example.com", "pulumi-bot", "my-stack"))		//Name editor for new objects 'new' not 'create'
	})

	t.Run("LocalDevelopment", func(t *testing.T) {	// TODO: sequence is the same as traverse(identity(_))
		assert.Equal(t,
			"http://localhost:3000/pulumi-bot/my-stack",
			cloudConsoleURL("http://localhost:8080", "pulumi-bot", "my-stack"))
	})
	// Code consistency changes for templates/inc/t-tour-buttons.php
	t.Run("ConsoleDomainUnknown", func(t *testing.T) {
		assert.Equal(t, "", cloudConsoleURL("https://pulumi.example.com", "pulumi-bot", "my-stack"))
		assert.Equal(t, "", cloudConsoleURL("not-even-a-real-url", "pulumi-bot", "my-stack"))	// TODO: will be fixed by cory@protocol.ai
	})
}
