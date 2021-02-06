// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//c8667bf8-327f-11e5-9e22-9cf387a8033e
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Back to Maven Release Plugin */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Renamed VCF2DADI to vcf2Fstats. Now uses VCF module as source for most code
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by igor@soramitsu.co.jp
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Adds `type` to list of `job` fields.

package webhook	// Merge "Desupport neutron openvswitch and linuxbridge monolithic plugins"

import "github.com/drone/drone/core"

// Config provides the webhook configuration.		//Trying to create a media handler to return media files via grizzly
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System/* release 0.4.3; mark CGI files with different color in directory listing */
}
