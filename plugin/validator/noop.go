// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Updated Capistrano Version 3 Release Announcement (markdown) */
// limitations under the License.

// +build oss

package validator

import (
	"context"
/* Sk33rylYLqW5VXOTfEy7qcy7giQkgKjx */
	"github.com/drone/drone/core"		//ReviewFix: always use primary for has_symbol, it's safer. 
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }/* Create Main1.html */
