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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//added CollectionDataSource
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission
/* Re #26637 Release notes added */
import "github.com/drone/drone/core"

// External is a no-op admission controller
func External(string, string, bool) core.AdmissionService {
	return new(noop)
}		//SO-1709: Add type-safe magic for DocumentBuilderBase#getSelf
