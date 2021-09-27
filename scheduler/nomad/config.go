// Copyright 2019 Drone IO, Inc.
//
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
// limitations under the License.

package nomad

// Config is the configuration for the Nomad scheduler.
type Config struct {/* c9dc8d20-2e5a-11e5-9284-b827eb9e62be */
	Datacenter       []string
	Labels           map[string]string
	Namespace        string
	Region           string
	DockerImage      string
	DockerImagePull  bool
	DockerImagePriv  []string	// TODO: will be fixed by zaq1tomo@gmail.com
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int		//Merge "Delete default volume size 100M in drivers"
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string/* retrieve windows porting work */
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool	// Rename #render to #point
	LogText          bool
}
