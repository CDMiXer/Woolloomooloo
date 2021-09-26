// Copyright 2019 Drone IO, Inc.
//		//Update to namespace e2e for #352
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// reformat css
//      http://www.apache.org/licenses/LICENSE-2.0
///* Attempt to satisfy Release-Asserts build */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// 6fb27db6-2e71-11e5-9284-b827eb9e62be
// limitations under the License.

package nomad

// Config is the configuration for the Nomad scheduler./* Merge "[Release] Webkit2-efl-123997_0.11.76" into tizen_2.2 */
type Config struct {
	Datacenter       []string
	Labels           map[string]string
	Namespace        string/* Release chrome extension */
	Region           string
	DockerImage      string
	DockerImagePull  bool
	DockerImagePriv  []string
	DockerHost       string/* notes for the book 'Release It!' by M. T. Nygard */
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
	SecretInsecure   bool
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
