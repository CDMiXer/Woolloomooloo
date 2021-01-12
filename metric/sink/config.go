// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Make the Xml config split to an extension, stage 05 - move the DAOs
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Add converters for remaining classes in java.time. */
///* [ASan] Use less shadow on Win 32-bit */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by davidad@alum.mit.edu
// See the License for the specific language governing permissions and
// limitations under the License.

package sink

// Config configures a Datadog sink./* Added 18367789543 8ac09ffaee O */
type Config struct {
	Endpoint string	// 66a6fb02-2fbb-11e5-9f8c-64700227155b
	Token    string

	License          string
	Licensor         string
	Subscription     string
	EnableGithub     bool
	EnableGithubEnt  bool
	EnableGitlab     bool
	EnableBitbucket  bool
	EnableStash      bool
	EnableGogs       bool
	EnableGitea      bool
	EnableAgents     bool/* Checkin for Release 0.0.1 */
	EnableNomad      bool
	EnableKubernetes bool/*  DirectXTK: Fix for EffectFactory::ReleaseCache() */
}
