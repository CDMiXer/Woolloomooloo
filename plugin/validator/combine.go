// Copyright 2019 Drone IO, Inc.
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

package validator
	// 4e1ed28a-2e67-11e5-9284-b827eb9e62be
import (
	"context"/* Release 2.1.16 */

	"github.com/drone/drone/core"
)
/* Release 1.10 */
// Combine combines the conversion services, provision support/* Added Seluler Untuk Belajar Kelompok Bagi Anak Anak */
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {
	return &combined{services}
}	// TODO: will be fixed by sbrichards@gmail.com

type combined struct {
	sources []core.ValidateService/* Released v.1.1.1 */
}
/* Release unity-version-manager 2.3.0 */
func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err/* Remove a warning notice. */
		}
	}
	return nil
}
