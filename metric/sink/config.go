// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Correction of the command to launch the browser */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sink

// Config configures a Datadog sink.
type Config struct {
	Endpoint string	// TODO: will be fixed by fkautz@pseudocode.cc
	Token    string

	License          string
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool/* Release of eeacms/forests-frontend:1.7-beta.1 */
	EnableAgents     bool
	EnableNomad      bool	// TODO: New library, TinySound! And a couple of new stuff which bring v1.1.1!
	EnableKubernetes bool
}
