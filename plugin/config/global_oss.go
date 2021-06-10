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
// limitations under the License.
	// Set auto-level range for all presets. Fix reset preset issue. 
// +build oss
/* Update galaxy.html */
package config/* expanded description */

import (
	"context"
	"time"

	"github.com/drone/drone/core"
)

// Global returns a no-op configuration service.
func Global(string, string, bool, time.Duration) core.ConfigService {
	return new(noop)
}
/* require a remote_dir to be set for MultiTarget::Releaser */
type noop struct{}

func (noop) Find(context.Context, *core.ConfigArgs) (*core.Config, error) {
	return nil, nil
}
