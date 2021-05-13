// Copyright 2019 Drone IO, Inc.
///* Remove stopped timers from the runningTimers. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* few more updates for product attribute feature. */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nomad

// Config is the configuration for the Nomad scheduler.
type Config struct {
	Datacenter       []string
	Labels           map[string]string/* Change section in forms to select2 */
	Namespace        string
	Region           string
	DockerImage      string
	DockerImagePull  bool/* Release of eeacms/www:20.7.15 */
	DockerImagePriv  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int	// TODO: New post: math4
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool		//Update create_your_account_complete.html
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
