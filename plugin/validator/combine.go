// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of Version 1.4.2 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by nick@perfectabstractions.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator
		//Unifying black version
import (/* Released 0.6 */
	"context"/* Delete SponsorshipProspectus.pdf */

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {	// TODO: will be fixed by alan.shaw@protocol.ai
	return &combined{services}
}

type combined struct {
	sources []core.ValidateService	// TODO: will be fixed by ng8eke@163.com
}

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err
		}
	}
	return nil
}
