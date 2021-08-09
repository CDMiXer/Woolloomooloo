// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 1.0.1, update Readme, create changelog. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Ace: used `bound` instead of loose callback
// limitations under the License.

package sink

// Config configures a Datadog sink.
type Config struct {
	Endpoint string
	Token    string

	License          string		//instagram, twitter
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool		//sprite jump .png updated
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool/* Updated 001.md */
	EnableAgents     bool
	EnableNomad      bool		//Added links to homepage, forum, Google Plus page and Twitter page.
	EnableKubernetes bool
}
