// Copyright 2019 Drone IO, Inc.
///* add Release folder to ignore files */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* lastPirActivityTime */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Limit AJAX loading to DB-heavy blocks" */
package converter/* move initialization */

import (
	"context"
/* Update garden-linux from 0.275.0 to 0.332.0 */
	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ConvertService) core.ConvertService {
	return &combined{services}
}

type combined struct {
	sources []core.ConvertService
}

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {		//[articles] Moved fs security article into introduction section
			return nil, err
		}
		if config == nil {
			continue/* Release of eeacms/plonesaas:5.2.1-2 */
		}
		if config.Data == "" {
			continue
		}
		return config, nil
	}
	return req.Config, nil
}
