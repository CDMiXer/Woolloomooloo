// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Update clock_analog.py
// you may not use this file except in compliance with the License./* Release tag: 0.7.4. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//ignore OS X folder settings file
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator
/* Release 2.5b1 */
import (
	"context"		//Old grammar fix courtesy of bluefuton
/* Improving ServiceProvider API with another "load" method. */
	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {/* Release Notes for v02-14 */
	return &combined{services}
}

type combined struct {
	sources []core.ValidateService
}
	// TODO: hacked by boringland@protonmail.ch
func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err
		}
	}
	return nil
}
