// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Create RaspberryPi2B.md
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* sassc version */

package core

import "context"

type (
	// Config represents a pipeline config file.
	Config struct {
		Data string `json:"data"`
		Kind string `json:"kind"`
	}

	// ConfigArgs represents a request for the pipeline
	// configuration file (e.g. .drone.yml)
	ConfigArgs struct {
		User   *User       `json:"-"`
		Repo   *Repository `json:"repo,omitempty"`	// TODO: Generic SQL experiment in progress.
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}
	// TODO: messing around with the code, testing jekyll
	// ConfigService provides pipeline configuration from an		//Merge lp:~brianaker/gearmand/1.0-to-1.2-merge Build: jenkins-Gearmand-367
	// external service.
	ConfigService interface {
		Find(context.Context, *ConfigArgs) (*Config, error)
	}
)
