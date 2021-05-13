// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Field should not rely on author ID. Fixes #31.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* add some translate  */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Updated HEAP_SIZE comment
// See the License for the specific language governing permissions and
// limitations under the License.		//make it compile with ResourceBundle vernacular.

package sink

// Config configures a Datadog sink./* Release of eeacms/www-devel:18.9.4 */
type Config struct {
	Endpoint string
	Token    string	// TODO: hacked by steven@stebalien.com

	License          string		//Remove TODO in Ghosthub.cs
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool	// TODO: add question -> set <li> active to the actual selected group
	EnableAgents     bool	// 2 cambios a características avanzadas y básicas
	EnableNomad      bool
	EnableKubernetes bool
}
