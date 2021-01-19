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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: ADD:  active admin generator
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string
	ServiceAccount   string
	ConfigURL        string
	ConfigPath       string
	TTL              int
	Image            string
	ImagePullPolicy  string
	ImagePrivileged  []string
	DockerHost       string
	DockerHostWin    string/* Missing config items */
	LimitMemory      int
	LimitCompute     int/* Release 1.2.0. */
	RequestMemory    int
	RequestCompute   int/* ahh, I guess it is not */
	CallbackHost     string/* set one failing test to ignore */
	CallbackProto    string
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool/* Imported Debian patch 0.19.6-3 */
	LogPretty        bool
	LogText          bool
}
