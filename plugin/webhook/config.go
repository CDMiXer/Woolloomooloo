// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Drop special PV handling for minishift
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release: Making ready for next release iteration 6.7.2 */
// limitations under the License.

package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration./* Activate all BUILD_ options if none was specified */
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System
}
