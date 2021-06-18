// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Updated TRS 80 (markdown)
// you may not use this file except in compliance with the License.	// TODO: will be fixed by ac0dem0nk3y@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//laravel4-turkish-documentation
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: 004188c8-2e76-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* test refactoring and new tests */
// limitations under the License./* [NOBTS] Fix duplicated scheduled run. */

package core

import "context"

type (
	// ConvertArgs represents a request to the pipeline
	// conversion service.
	ConvertArgs struct {
		User   *User       `json:"-"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}

	// ConvertService converts non-native pipeline
	// configuration formats to native configuration
	// formats (e.g. jsonnet to yaml).
	ConvertService interface {
		Convert(context.Context, *ConvertArgs) (*Config, error)
	}
)
