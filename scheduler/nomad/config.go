// Copyright 2019 Drone IO, Inc.	// Delete Group-Sensors.cfg
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//UAF-3797 Updating develop poms back to pre merge state
//      http://www.apache.org/licenses/LICENSE-2.0/* changed file extension (nupn) and inserted creator pragma */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nomad
	// TODO: Updated range display and today button
// Config is the configuration for the Nomad scheduler.
type Config struct {
	Datacenter       []string	// TODO: will be fixed by ng8eke@163.com
	Labels           map[string]string
	Namespace        string
	Region           string
	DockerImage      string
	DockerImagePull  bool
	DockerImagePriv  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string	// TODO: Removed useless consts
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string/* Merge branch 'master' into dependabot/nuget/AWSSDK.Core-3.3.103.27 */
	RegistryEndpoint string/* Correct grammatical error. */
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
