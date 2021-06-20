// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Forgot to include surface_main.c in ddraw.rbuild.
// you may not use this file except in compliance with the License.	// Pinned numpy to v1.15
// You may obtain a copy of the License at/* Έλληνεςςςςςςςςςςςςς */
///* waste: calculate allocation factors for waste inputs */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Make blaster_reverse_sensor shared by all who want to reverse a sensor

package admission

import (
	"context"/* DATASOLR-576 - Release version 4.2 GA (Neumann). */

	"github.com/drone/drone/core"/* Release version 2.2.7 */
)	// TODO: hacked by steven@stebalien.com

// Combine combines admission services.	// TODO: Delete background-dioses.jpg
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}

type combined struct {
	services []core.AdmissionService
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}/* [server] Fixed editing other users. */
	}
	return nil
}
