// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Merge "Implement "IPAllocation" router ports allocated retrieval"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* [REF] use single implementation for name_search of Country and CountryState */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by alan.shaw@protocol.ai
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// Merge branch 'master' into ZON-4012
/* Merge "Set http_proxy to retrieve the signed Release file" */
// +build oss

package converter

import (
	"github.com/drone/drone/core"
)/* Release Notes for v02-08 */
/* decoder/dsf: use the block count internally */
// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return new(noop)
}
