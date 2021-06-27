// Copyright 2019 Drone IO, Inc.	// TODO: hacked by martin2cai@hotmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: I made a stupid mistake.
// You may obtain a copy of the License at/* Release new version 2.5.9: Turn on new webRequest code for all Chrome 17 users */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// definition of pp routine: EquidistanceProvider
// limitations under the License.

package webhook	// TODO: adapted desc to current implementation

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string/* create main_css */
	System   *core.System
}
