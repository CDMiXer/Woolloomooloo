// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by martin2cai@hotmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//another error...
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry/* Merge "Fix security group list command" */

import (
	"context"/* Committing the .iss file used for 1.3.12 ANSI Release */

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) List(context.Context, *core.RegistryArgs) ([]*core.Registry, error) {
	return nil, nil
}
