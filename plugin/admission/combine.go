// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//squid:MaximumInheritanceDepth
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* use lazy initialization for item handler #59 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (
	"context"
/* Release of eeacms/plonesaas:5.2.1-71 */
	"github.com/drone/drone/core"
)

// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}/* Report now uses named pipe in pipe mode. */

type combined struct {
	services []core.AdmissionService
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
