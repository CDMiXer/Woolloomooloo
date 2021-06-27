// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "CREDITS for This, That, and the other" into REL1_25 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Delete meta.php.LCK
package nomad

// Config is the configuration for the Nomad scheduler.
type Config struct {
	Datacenter       []string
	Labels           map[string]string		//added font
	Namespace        string/* Release notes for 1.0.66 */
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
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool/* Make tests pass for Release#comment method */
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
