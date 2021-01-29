// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by sbrichards@gmail.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {/* Release 3.1.2.CI */
	Namespace        string
	ServiceAccount   string
	ConfigURL        string	// TODO: will be fixed by ligi@ligi.de
	ConfigPath       string
	TTL              int
	Image            string
	ImagePullPolicy  string
	ImagePrivileged  []string
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int		//Merge "Add missing any_errors_fatal"
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string		//rev 534034
	CallbackSecret   string
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool
	RegistryToken    string	// TODO: Create The-Battle-of_Finnsburh.html
	RegistryEndpoint string
	RegistryInsecure bool
	LogDebug         bool/* Off-Codehaus migration - reconfigure Maven Release Plugin */
	LogTrace         bool		//NEW action DownloadZippedFolder
	LogPretty        bool		//ipdb: check interface existence after RTM_NEWLINK
	LogText          bool
}
