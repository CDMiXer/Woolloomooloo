// Copyright 2019 Drone IO, Inc./* #195 fix url and click selector */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Add some Release Notes for upcoming version */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update URLHelper.php */
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by witek@enjin.io
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sink

// Config configures a Datadog sink.
type Config struct {		//add sqrt, rational_number, e2_1
	Endpoint string		//Renamed LCOMPreferencePage to CohesionPreferencePage
	Token    string
		//#12 font-size: 14px;  #126 @media(min-width:650);
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
