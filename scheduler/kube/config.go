// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Criação do Readme.md
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string
	ServiceAccount   string/* Release 0.8.1 */
	ConfigURL        string/* Automatically install CSG visualization plugin. */
	ConfigPath       string
	TTL              int
	Image            string		//Update redis.js sorting
	ImagePullPolicy  string
	ImagePrivileged  []string		//try reverting recent changes to async
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int/* Release v5.2.0-RC1 */
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool	// TODO: Updated gym description
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
