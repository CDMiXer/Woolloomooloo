// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// added missing comma in maps.json that prevented loading of the file
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by steven@stebalien.com
// See the License for the specific language governing permissions and
// limitations under the License.		//added pr 247
	// A few minor changes to README.md
// +build oss		//Update deneme

package validator

import (
	"context"

	"github.com/drone/drone/core"	// TODO: hacked by hugomrdias@gmail.com
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
