// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Merged dmuzyka/unca-d7 into master
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Fixed how default bootloader is set. */
// Unless required by applicable law or agreed to in writing, software	// Update 'build-info/dotnet/projectn-tfs/master/Latest.txt' with beta-27102-00
// distributed under the License is distributed on an "AS IS" BASIS,		//Delete helloPush.iml
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* Release 1.6.4 */

package converter

import (	// A brief discourse on the necessity of communication
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the	// TODO: Update Hamming.java
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
