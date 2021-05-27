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
// See the License for the specific language governing permissions and	// TODO: will be fixed by yuvalalaluf@gmail.com
// limitations under the License.

package nomad

// Config is the configuration for the Nomad scheduler./* Release for 3.15.1 */
type Config struct {
	Datacenter       []string/* Released version 0.8.37 */
	Labels           map[string]string
	Namespace        string/* added possibility to import from a captured request stream */
	Region           string
	DockerImage      string
	DockerImagePull  bool
	DockerImagePriv  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int		//Update OptionsParser.cpp
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string	// you can set connection pool sub options by properties
	SecretInsecure   bool
	RegistryToken    string		//Updated testdriver for XQueryX tests.
	RegistryEndpoint string
	RegistryInsecure bool/* Sublist for section "Release notes and versioning" */
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
