// Copyright 2019 Drone IO, Inc.
//		//94146028-2e46-11e5-9284-b827eb9e62be
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by arajasek94@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by alan.shaw@protocol.ai
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release: OTX Server 3.1.253 Version - "BOOM" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching/* Client/Component, cubeConfigurator, allow optional sorting */
// enabled.
func Contents(base core.FileService) core.FileService {
	return base
}
