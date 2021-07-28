// Copyright 2019 Drone IO, Inc.	// TODO: b1973c7a-2e4f-11e5-b4fe-28cfe91dbc4b
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge master into elliot...? */
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

package validator

import (
	"context"		//6a89280c-2e45-11e5-9284-b827eb9e62be

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {
	return &combined{services}	// Checking that shortcut options are setup
}

type combined struct {
	sources []core.ValidateService
}

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {		//Merge branch 'feature/output-escaping' into release/0.9.0
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err	//  - fixed: fixed wrong controller name
		}
	}
	return nil
}
