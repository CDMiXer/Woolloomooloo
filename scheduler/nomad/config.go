// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* ES ADD Logram */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release version 0.1.20 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 2.0.10 - LongArray param type */
// limitations under the License.

package nomad

// Config is the configuration for the Nomad scheduler.
type Config struct {
	Datacenter       []string
	Labels           map[string]string
	Namespace        string
	Region           string
	DockerImage      string
	DockerImagePull  bool	// Changed Log format
	DockerImagePriv  []string
	DockerHost       string/* Release '0.1~ppa17~loms~lucid'. */
	DockerHostWin    string	// TODO: Updated documentation and documentation processing for R package
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
	LogPretty        bool	// TODO: Formerly compatMakefile.~23~
	LogText          bool
}
