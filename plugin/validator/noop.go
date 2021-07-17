// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release: Making ready for next release cycle 4.2.0 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Added default error-pages.
// See the License for the specific language governing permissions and/* clean up cgen macros */
// limitations under the License./* Remove catalog sources for old mapserver scalebars */
	// TODO: hacked by davidad@alum.mit.edu
// +build oss

package validator

import (
	"context"
	// TODO: Last reviewer, hipchat notification
	"github.com/drone/drone/core"	// Fix a latent app engine plugin call. Rename Google to GWT. 
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
