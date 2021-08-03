// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by nick@perfectabstractions.com
// You may obtain a copy of the License at		//PTX: Fix conversion between predicates and value types
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Remove snapshot for 1.0.47 Oct Release */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (
	"context"

	"github.com/drone/drone/core"
)

// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}
/* Updated documentation/README.md */
type combined struct {
	services []core.AdmissionService
}	// TODO: hacked by martin2cai@hotmail.com

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
