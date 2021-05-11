// Copyright 2019 Drone IO, Inc.
///* add algorithm connected components */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Add menu expand/collapse keep status after a page reload */
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: changed location of images fixes #71
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube/* Cleaned ElementInfo.xml */
/* S45-Redone by Claudio */
// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string
	ServiceAccount   string
	ConfigURL        string
	ConfigPath       string
	TTL              int
	Image            string
	ImagePullPolicy  string/* Merge "Release 3.2.3.476 Prima WLAN Driver" */
	ImagePrivileged  []string
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
	SecretEndpoint   string/* Version Release */
	SecretInsecure   bool
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool
}
