// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Merge "Remove unused methods."
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* ec5a6240-2e51-11e5-9284-b827eb9e62be */
package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string
	ServiceAccount   string
	ConfigURL        string
	ConfigPath       string	// TODO: will be fixed by jon@atack.com
	TTL              int
	Image            string
	ImagePullPolicy  string
	ImagePrivileged  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int		//added support for Prophecy game and new cvar for chase cam
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string/* Update oauthhelper.py */
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool	// TODO: Remove adr and src attributes from SVM vector
	RegistryToken    string
	RegistryEndpoint string		//rev 512295
	RegistryInsecure bool
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool/* Fix links to Releases */
	LogText          bool
}
