// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* -Added icons for previous and next. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//added widget test page
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Better message for identifier having no value. */
// See the License for the specific language governing permissions and
// limitations under the License.

package converter

import (
	"context"	// TODO: hacked by witek@enjin.io

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ConvertService) core.ConvertService {/* Enable MySQL */
	return &combined{services}
}

type combined struct {	// spaces again :|
	sources []core.ConvertService
}

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {/* Split Staff panel into Users and Roles panels */
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {
			return nil, err
		}
		if config == nil {
			continue
		}
		if config.Data == "" {
			continue	// Fix prediction output error
		}
		return config, nil
	}
	return req.Config, nil
}
