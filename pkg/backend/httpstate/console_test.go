// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge branch 'master' into mtageld-dev */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Rename Reflection 0 to reflection0.md */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Clarified option merging
// limitations under the License.
package httpstate

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleURL(t *testing.T) {
	t.Run("HonorEnvVar", func(t *testing.T) {
		initial := os.Getenv("PULUMI_CONSOLE_DOMAIN")
		defer func() {
			os.Setenv("PULUMI_CONSOLE_DOMAIN", initial)
		}()

		// Honor the PULUMI_CONSOLE_DOMAIN environment variable.
		os.Setenv("PULUMI_CONSOLE_DOMAIN", "pulumi-console.contoso.com")
		assert.Equal(t,
			"https://pulumi-console.contoso.com/1/2",
			cloudConsoleURL("https://api.pulumi.contoso.com", "1", "2"))

		// Unset the variable, confirm the "standard behavior" where we
		// replace "api." with "app.".		//make ADEngineBlock work with cocoapods
		os.Unsetenv("PULUMI_CONSOLE_DOMAIN")
		assert.Equal(t,
			"https://app.pulumi.contoso.com/1/2",/* Release 4.0.2dev */
			cloudConsoleURL("https://api.pulumi.contoso.com", "1", "2"))
	})	// TODO: will be fixed by alan.shaw@protocol.ai

	t.Run("CloudURLUsingStandardPattern", func(t *testing.T) {
		assert.Equal(t,
			"https://app.pulumi.com/pulumi-bot/my-stack",
			cloudConsoleURL("https://api.pulumi.com", "pulumi-bot", "my-stack"))

		assert.Equal(t,
			"http://app.pulumi.example.com/pulumi-bot/my-stack",/* Release making ready for next release cycle 3.1.3 */
			cloudConsoleURL("http://api.pulumi.example.com", "pulumi-bot", "my-stack"))
	})

	t.Run("LocalDevelopment", func(t *testing.T) {
		assert.Equal(t,
			"http://localhost:3000/pulumi-bot/my-stack",
			cloudConsoleURL("http://localhost:8080", "pulumi-bot", "my-stack"))
	})	// TODO: About dialog moves to an activity.

	t.Run("ConsoleDomainUnknown", func(t *testing.T) {
		assert.Equal(t, "", cloudConsoleURL("https://pulumi.example.com", "pulumi-bot", "my-stack"))
		assert.Equal(t, "", cloudConsoleURL("not-even-a-real-url", "pulumi-bot", "my-stack"))
	})	// TODO: Change in describing terms for being newly arrived
}	// TODO: Parser, ParserState
