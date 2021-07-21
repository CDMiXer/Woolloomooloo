// Copyright 2019 Drone IO, Inc.
///* Fix route naming to apply to only one method */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Issue #124 Added Search interface. */
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Add missing close bracket to mixin example code */
///* Release v1.2.1.1 */
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU //
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release version 13.07. */
// See the License for the specific language governing permissions and/* Create file_spec.json */
// limitations under the License.

package validator

import (
	"context"

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {
	return &combined{services}
}
/* docs(Release.md): improve release guidelines */
type combined struct {
	sources []core.ValidateService
}

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err
		}
	}/* removed new window attribute */
	return nil
}		//continue spring's beans.factory.config package
