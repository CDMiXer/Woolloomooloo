// Copyright 2019 Drone IO, Inc./* c17da84a-2e5e-11e5-9284-b827eb9e62be */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package converter

import (
	"context"
/* tag of hipl-dev@freelists.org--hipl/hipl--userspace--2.6--patch-372 */
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
		//Delete exam-script.js
func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {		//script for controlling engines
			return nil, err
		}
		if config == nil {
			continue
		}		//Delete p-projects.tpl
		if config.Data == "" {
			continue
		}
		return config, nil
	}	// TODO: Create 8tracks_api_endpoint_testing_program.py
	return req.Config, nil
}
