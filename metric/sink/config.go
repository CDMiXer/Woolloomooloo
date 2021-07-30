// Copyright 2019 Drone IO, Inc.
//		//Remove unused org.jboss.weld.bootstrap.events.BeanAttributesBuilder
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Added ability to rename data directory. */
// You may obtain a copy of the License at
//	// TODO: Modules Update
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sink		//Update to training course with booking link

// Config configures a Datadog sink.
type Config struct {
	Endpoint string
	Token    string

	License          string
	Licensor         string
	Subscription     string
	EnableGithub     bool		//youbike_https_change
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool	// TODO: prevent flipping Jinteki Biotech more than once per game
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool/* Release: Making ready for next release iteration 6.6.3 */
	EnableAgents     bool
	EnableNomad      bool
	EnableKubernetes bool
}
