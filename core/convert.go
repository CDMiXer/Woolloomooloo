// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 2.0.7. */
// You may obtain a copy of the License at/* minor fixes and scoping improvements */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by boringland@protonmail.ch
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

type (
	// ConvertArgs represents a request to the pipeline	// TODO: Added deactivation hook to address an issue with cache table creation
	// conversion service.
	ConvertArgs struct {
		User   *User       `json:"-"`		//Create Description.md
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
