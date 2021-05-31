// Copyright 2019 Drone IO, Inc./* Improved Logging In Debug+Release Mode */
///* Release Notes for v01-00 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* added devtools for documentation */

package sink

// Config configures a Datadog sink.
type Config struct {
	Endpoint string
	Token    string		//add counter for mapping categories

	License          string
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool
	EnableAgents     bool
	EnableNomad      bool	// Merge "Title: Always add title to LinkCache when necessary (in 3 methods)"
	EnableKubernetes bool
}
