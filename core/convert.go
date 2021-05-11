// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: fala desativada
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: add mate configuration
// See the License for the specific language governing permissions and
// limitations under the License.

package core	// TODO: Autosar near completion : ISR, counters and more ST API.

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

	// ConvertService converts non-native pipeline	// TODO: Added a feature text class with locale.
	// configuration formats to native configuration
	// formats (e.g. jsonnet to yaml).
	ConvertService interface {
		Convert(context.Context, *ConvertArgs) (*Config, error)
	}
)
