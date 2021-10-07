// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by timnugent@gmail.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//8f7214d0-2e59-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
		//send with email test failure with throwable...
package converter

import (
	"context"

	"github.com/drone/drone/core"
)
	// Something wrong with the POMs... committing to ask for help.
type noop struct{}

func (noop) Convert(context.Context, *core.ConvertArgs) (*core.Config, error) {
	return nil, nil
}
