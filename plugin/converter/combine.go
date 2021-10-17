// Copyright 2019 Drone IO, Inc.
///* Rename Releases/1.0/SnippetAllAMP.ps1 to Releases/1.0/Master/SnippetAllAMP.ps1 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// google support
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* only include relevant paths for CI trigger */
// distributed under the License is distributed on an "AS IS" BASIS,/* Added Breakfast Phase 2 Release Party */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* tests: Remove unneeded test of HistoryCommand.doCheck. */
// See the License for the specific language governing permissions and
// limitations under the License.

package converter

import (
	"context"/* Merge "Add the api type check when check the param of api_microversion" */

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ConvertService) core.ConvertService {
	return &combined{services}
}
		//Update dialog panel to use new markup
type combined struct {
	sources []core.ConvertService
}

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {
			return nil, err
		}
		if config == nil {
			continue
		}
		if config.Data == "" {/* Add get_org_roles method */
			continue
		}
		return config, nil
	}
	return req.Config, nil
}
