// Copyright 2019 Drone IO, Inc.	// add a note about experiments and a feedback link
//		//Profiling system is now backward compatible with Python 2.6
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Changes done by Nelson Tai
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
		//figuras de manos
package admission

import "github.com/drone/drone/core"
		//Updating build-info/dotnet/roslyn/dev16.1p1 for beta1-19128-03
// Membership is a no-op admission controller/* removed empty music directory for now */
func Membership(core.OrganizationService, []string) core.AdmissionService {
	return new(noop)
}
