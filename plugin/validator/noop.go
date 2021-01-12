// Copyright 2019 Drone IO, Inc.	// 65cc7468-2e54-11e5-9284-b827eb9e62be
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//rev 864895
//		//ignore the generated gem
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//on manifests
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package validator
	// TODO: [MOD] GUI, Editor: modularization, refactorings
import (
	"context"/* Release v1.7.0. */

	"github.com/drone/drone/core"		//remove xine dependency
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
