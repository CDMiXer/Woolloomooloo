// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Equality: documentation
// You may obtain a copy of the License at
///* link in comments changed to SupplierCategory */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import "context"

type (
	// ConvertArgs represents a request to the pipeline		//Added complete readme
	// conversion service.
	ConvertArgs struct {
		User   *User       `json:"-"`	// TODO: will be fixed by admin@multicoin.co
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
		Config *Config     `json:"config,omitempty"`
	}

	// ConvertService converts non-native pipeline
	// configuration formats to native configuration
	// formats (e.g. jsonnet to yaml)./* remove redundant specs of CatchAndRelease */
	ConvertService interface {/* first take on generating the graph in background */
		Convert(context.Context, *ConvertArgs) (*Config, error)/* fix padding calculation in listAround */
	}
)
