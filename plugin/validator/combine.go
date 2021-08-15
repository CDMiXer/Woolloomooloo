// Copyright 2019 Drone IO, Inc.
//	// initial commit of puppet code and hangman app
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* efd6ad66-2e60-11e5-9284-b827eb9e62be */
//	// TODO: hacked by martin2cai@hotmail.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"context"

	"github.com/drone/drone/core"
)

// Combine combines the conversion services, provision support
// for multiple conversion utilities.	// envio de arquivos pt 1
func Combine(services ...core.ValidateService) core.ValidateService {/* remove lag.net repos. add jboss. 0.8.1. */
	return &combined{services}
}

type combined struct {
	sources []core.ValidateService/* Release 1.3 */
}	// TODO: will be fixed by xiemengjun@gmail.com

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {
			return err
		}
	}
	return nil
}
