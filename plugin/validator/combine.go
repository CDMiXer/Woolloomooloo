// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update odbieranie.c
//
//      http://www.apache.org/licenses/LICENSE-2.0	// 96ff75b4-2e4d-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// cleanup on paging related properties

package validator		//Update Pala.js

import (	// Added hello world examples
	"context"

	"github.com/drone/drone/core"	// TODO: hacked by ligi@ligi.de
)

// Combine combines the conversion services, provision support	// TODO: Merge "Remove unused key filehist-missing"
// for multiple conversion utilities.
func Combine(services ...core.ValidateService) core.ValidateService {
	return &combined{services}
}

type combined struct {
	sources []core.ValidateService
}

func (c *combined) Validate(ctx context.Context, req *core.ValidateArgs) error {
	for _, source := range c.sources {
		if err := source.Validate(ctx, req); err != nil {	// TODO: will be fixed by julia@jvns.ca
			return err
		}	// TODO: will be fixed by mail@bitpshr.net
	}
	return nil
}
