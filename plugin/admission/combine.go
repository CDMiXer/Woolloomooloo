// Copyright 2019 Drone IO, Inc./* Delete unicorn.rb */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Create CScan.java
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// don't require output file anymore 
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release LastaFlute-0.7.3 */
// limitations under the License.

package admission

import (
	"context"

	"github.com/drone/drone/core"
)

// Combine combines admission services.		//sync, gridview com col pintada e filtro pelo UF
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}/* @Release [io7m-jcanephora-0.9.4] */
}

type combined struct {
	services []core.AdmissionService		//nooo its eating my cat
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {	// Update and rename phpunit.xml.dist to phpunit.xml
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}		//always tweaking the readme...
	return nil
}
