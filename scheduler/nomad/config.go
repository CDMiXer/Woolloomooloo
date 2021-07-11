// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Create 1.0_Final_ReleaseNote.md */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Add two beautiful unsplash photos */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by magik6k@gmail.com
package nomad

// Config is the configuration for the Nomad scheduler.
type Config struct {
	Datacenter       []string
	Labels           map[string]string
	Namespace        string
	Region           string
	DockerImage      string	// de41e4b8-2e4b-11e5-9284-b827eb9e62be
	DockerImagePull  bool
	DockerImagePriv  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int
	CallbackHost     string/* Merge "Remove unused domain config method paramters" */
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
