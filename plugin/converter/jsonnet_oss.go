// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Dev checkin #324.  Fix export on individual flow.
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Fix problem where row numbers in selections are incorrect after a sort

// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return new(noop)/* Release of eeacms/eprtr-frontend:2.0.5 */
}
