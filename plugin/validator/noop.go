// Copyright 2019 Drone IO, Inc./* Release ntoes update. */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Rename code.sh to raiSah6ashiraiSah6ashiraiSah6ashi.sh */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: First call resizeMedia before changing width attr
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Copy nested fragments
// limitations under the License.

// +build oss

package validator

import (
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }/* Released MotionBundler v0.1.1 */
