// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* HAP-41 - Include a link to the builds list to jobs that use Artifactory plugin */
///* Added author contribution to about dialog for restart icon. */
// Unless required by applicable law or agreed to in writing, software		//Mention FB port in Readme
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: replace with more modern word
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package validator
	// Create tech-4-finance.md
import (
	"context"
/* Fix commited regressions still block CI, They must be FIx Released to unblock */
	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
