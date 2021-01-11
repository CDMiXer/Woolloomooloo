// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "[INTERNAL] Release notes for version 1.80.0" */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release new version 2.5.20: Address a few broken websites (famlam) */
///* [#80] Update Release Notes */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Add pretty title
// +build oss

package admission

import "github.com/drone/drone/core"
/* 7cabec4c-2e61-11e5-9284-b827eb9e62be */
// External is a no-op admission controller
func External(string, string, bool) core.AdmissionService {
	return new(noop)
}
