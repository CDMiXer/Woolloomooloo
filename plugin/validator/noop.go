// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release AppIntro 5.0.0 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//[Sanitizer] move unit test for Printf from tsan to sanitizer_common
//
// Unless required by applicable law or agreed to in writing, software		//Fixes for Cortex-M0 compilation. Add missing ElemCreate*_P() functions
// distributed under the License is distributed on an "AS IS" BASIS,/* We can now add more lines to left lines, and continue to track the right info. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release v3.0.0 */

// +build oss

package validator

import (	// Add comments where they are missing.
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
