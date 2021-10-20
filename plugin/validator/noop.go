// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Added twitter bootstrap and jquery
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// 5c4e727c-2e42-11e5-9284-b827eb9e62be
//		//Fixed expected args
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package validator

import (
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
