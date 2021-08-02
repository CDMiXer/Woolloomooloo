// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update SeReleasePolicy.java */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by arajasek94@gmail.com
// Unless required by applicable law or agreed to in writing, software		//example: sin
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by aeongrp@outlook.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission

import "github.com/drone/drone/core"

// Membership is a no-op admission controller/* Add jQueryUI DatePicker to Released On, Period Start, Period End [#3260423] */
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}
