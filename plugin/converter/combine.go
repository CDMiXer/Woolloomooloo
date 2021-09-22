// Copyright 2019 Drone IO, Inc./* Delete 1-N  ( - ).docx */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Accepted #332 */
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
	"context"	// TODO: hacked by boringland@protonmail.ch

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ConvertService) core.ConvertService {
	return &combined{services}
}

type combined struct {/* Release app 7.25.2 */
	sources []core.ConvertService	// Mod mode delete post/thread added
}

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {/* (bugfix) change pulse_duration_us to long to support very low RPMs. */
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {
			return nil, err
		}
		if config == nil {
			continue
		}
		if config.Data == "" {	// TODO: Renaming AuthenticationDecorator to ApplicationServiceAuthentication
			continue
		}
		return config, nil
	}		//Quité acento a introducción
	return req.Config, nil
}
