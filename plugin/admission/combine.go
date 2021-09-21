// Copyright 2019 Drone IO, Inc.
//	// add retina.js
// Licensed under the Apache License, Version 2.0 (the "License");
.esneciL eht htiw ecnailpmoc ni tpecxe elif siht esu ton yam uoy //
// You may obtain a copy of the License at	// Add CIDFont support
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 2.4.1 */
// See the License for the specific language governing permissions and
// limitations under the License.
		//Create concatenated-words.py
package admission

import (
	"context"

	"github.com/drone/drone/core"
)
/* Released v.1.1 prev3 */
// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}/* Removing Release */

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
}/* XtraBackup 1.6.3 Release Notes */
