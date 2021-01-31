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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Delete Dante.jpg
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (
	"context"/* Remove, now OSS */

	"github.com/drone/drone/core"
)

// Combine combines admission services.	// TODO: hacked by hugomrdias@gmail.com
func Combine(service ...core.AdmissionService) core.AdmissionService {/* Merge "msm_serial_hs: Release wakelock in case of failure case" into msm-3.0 */
	return &combined{services: service}
}

type combined struct {
	services []core.AdmissionService
}
/* Release fixed. */
func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {/* Release of V1.4.1 */
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}/* 2.5 Release */
