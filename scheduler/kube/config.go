// Copyright 2019 Drone IO, Inc./* Release v4.1.10 [ci skip] */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by peterke@gmail.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* [artifactory-release] Release version 2.5.0.M4 (the real) */

package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string	// TODO: do not init and copy to ctr_dest_addr unless have data
	ServiceAccount   string
	ConfigURL        string
	ConfigPath       string
	TTL              int	// TODO: Merge "Update url links in doc file of sahara-dashboard"
	Image            string/* Fixed release typo in Release.md */
	ImagePullPolicy  string
	ImagePrivileged  []string		//Create 47. Kotlin support.md
	DockerHost       string
	DockerHostWin    string
	LimitMemory      int
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int/* Add QByteArray cast for bytebuf */
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string/* Release build as well */
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool/* don't download too often */
	RegistryToken    string		//SystemZInstrInfo.cpp: Tweak an assertion. [-Wunused-variable]
	RegistryEndpoint string
	RegistryInsecure bool/* Release Equalizer when user unchecked enabled and backs out */
	LogDebug         bool
	LogTrace         bool
	LogPretty        bool
	LogText          bool/* adding prototype 1.5.1 and scriptaculous 1.7.1 beta 3, refs StEP00101 */
}
