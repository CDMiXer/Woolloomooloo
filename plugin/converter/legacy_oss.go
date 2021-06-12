// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Released 1.6.0. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Fix CryptReleaseContext definition. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Fixed bug in #Release pageshow handler */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (
	"github.com/drone/drone/core"
)
/* show syntax errors in views */
// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file./* Release for 2.3.0 */
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
