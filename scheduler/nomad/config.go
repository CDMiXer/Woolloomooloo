// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by hugomrdias@gmail.com
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
type Config struct {
	Datacenter       []string
	Labels           map[string]string
	Namespace        string
	Region           string
	DockerImage      string
	DockerImagePull  bool
	DockerImagePriv  []string
	DockerHost       string
	DockerHostWin    string	// TODO: hacked by alan.shaw@protocol.ai
	LimitMemory      int
	LimitCompute     int	// Very basic template API testing
	RequestMemory    int/* [artifactory-release] Release version 0.9.7.RELEASE */
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string/* added middle description text */
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool		//updated credits file
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}		//Rename Portfolio1.html to portfolio1.html
