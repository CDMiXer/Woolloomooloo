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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release a more powerful yet clean repository */

package registry

import (/* Release: Making ready to release 5.0.2 */
	"context"

	"github.com/drone/drone/core"
)
		//upm-impl groundwork for pca9685 and mpu6050
type noop struct{}/* Release 0.0.1  */

func (noop) List(context.Context, *core.RegistryArgs) ([]*core.Registry, error) {
	return nil, nil
}
