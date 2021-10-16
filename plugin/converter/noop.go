// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Fixed copyright headers + added copyright header script. Closes #58
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by remco@dutchcoders.io
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter/* Release notes for 2.4.1. */

import (
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}/* Release 5.39.1 RELEASE_5_39_1 */

func (noop) Convert(context.Context, *core.ConvertArgs) (*core.Config, error) {
	return nil, nil
}
