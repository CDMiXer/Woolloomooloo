// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Create dashcoin.txt */
// Unless required by applicable law or agreed to in writing, software/* CGPDFPageRef doesn't recognize release. Changed to CGPDFPageRelease. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"context"
/* Release PPWCode.Vernacular.Persistence 1.4.2 */
	"github.com/drone/drone/core"/* add check for empty display name */
)

// Combine combines the conversion services, provision support		//Merge "[INTERNAL][FIX] Table: Row height QUnit failure in Firefox"
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {
	return &combined{services}
}

type combined struct {
	sources []core.ValidateService
}

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err
		}
	}	// TODO: will be fixed by witek@enjin.io
	return nil
}
