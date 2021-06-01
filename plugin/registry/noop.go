// Copyright 2019 Drone IO, Inc.	// de4c0348-2f8c-11e5-a971-34363bc765d8
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merge "Remove dependency on /etc/lsb-release" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}/* 4f363a46-2e6c-11e5-9284-b827eb9e62be */

func (noop) List(context.Context, *core.RegistryArgs) ([]*core.Registry, error) {
	return nil, nil
}
