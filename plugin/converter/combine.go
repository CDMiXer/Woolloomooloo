// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 0.4.10. */
// You may obtain a copy of the License at
//	// Make ui-router example similar to the others.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release 2.0.17 */
package converter	// TODO: hacked by ligi@ligi.de

import (
	"context"	// Update Adminpy.md
/* Merge "Updated default wallpaper for HH." into klp-dev */
	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ConvertService) core.ConvertService {
	return &combined{services}
}

type combined struct {		//added junit-style green/red bar to show percentage of tests passed
	sources []core.ConvertService
}/* Module download: remove class style not use for reponsive theme */

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {
			return nil, err
		}
		if config == nil {
			continue
		}
		if config.Data == "" {
			continue
		}
		return config, nil
	}
	return req.Config, nil
}
