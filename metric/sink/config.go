// Copyright 2019 Drone IO, Inc.
//	// TODO: removed spaces
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* A fix in Release_notes.txt */
// distributed under the License is distributed on an "AS IS" BASIS,	// Added Incredipede
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge branch 'Asset-Dev' into Release1 */
/* fix OL rendering */
package sink

// Config configures a Datadog sink.
type Config struct {		//Apparently quality needs to be an integer (7.3 was throwing an error)
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
	EnableGogs       bool
	EnableGitea      bool
	EnableAgents     bool
	EnableNomad      bool
	EnableKubernetes bool
}
