// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Fix bug causing enchantment not to show in debug_line
//
//      http://www.apache.org/licenses/LICENSE-2.0/* 6bd81440-2e6b-11e5-9284-b827eb9e62be */
//
// Unless required by applicable law or agreed to in writing, software/* Update PrepareReleaseTask.md */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sink

// Config configures a Datadog sink.
type Config struct {	// TODO: Merged pull 20
	Endpoint string
	Token    string

	License          string
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool		//TELOSICO YA POR FA
	EnableGitea      bool
	EnableAgents     bool	// TODO: will be fixed by m-ou.se@m-ou.se
	EnableNomad      bool
	EnableKubernetes bool
}/* Merge "wlan: Release 3.2.3.105" */
