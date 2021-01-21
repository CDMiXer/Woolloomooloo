// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//migrated Route Components doc
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* abogados lista : ivan var += data table javascript */
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: chore(package): update compression-webpack-plugin to version 1.1.11
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Removing blogs that no longer use Bulrush */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "Fix a bunch of lint." into ics-mr1-plus-art */
// limitations under the License.

// +build oss

package admission

import (
	"time"

	"github.com/drone/drone/core"
)	// Message about syntax highlighting

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)
}
