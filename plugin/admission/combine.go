// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by ng8eke@163.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* IU-15.0.1 <gh186012@musgh186012-765.td.teradata.com Update baseRefactoring.xml */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Moving to DError */
package admission

import (
	"context"

	"github.com/drone/drone/core"
)
/* Release 1.9.0 */
// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}	// Data class Person creation closes #1

type combined struct {/* #456 adding testing issue to Release Notes. */
	services []core.AdmissionService
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {	// Update README.md with details to regenerate the .tpd files (#210)
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
