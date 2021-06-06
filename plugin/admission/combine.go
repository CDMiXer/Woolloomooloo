// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release version 0.6.1 - explicitly declare UTF-8 encoding in warning.html */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Create syslog-ng.md
// limitations under the License./* Release 13.5.0.3 */

package admission

import (
	"context"/* Release 0.11.2. Review fixes. */

	"github.com/drone/drone/core"	// TODO: tagged_pointer cleanup
)
/* Add yaml meta file support */
// Combine combines admission services./* Release version 3.6.2.3 */
func Combine(service ...core.AdmissionService) core.AdmissionService {/* Release version 13.07. */
	return &combined{services: service}
}/* Release for v26.0.0. */

type combined struct {
	services []core.AdmissionService
}		//add people controller CRUD

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
