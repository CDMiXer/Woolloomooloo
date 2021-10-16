// Copyright 2019 Drone IO, Inc.
///* Preparing for 0.1.5 Release. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release notes 7.1.11 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* chore: add dry-run option to Release workflow */
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"testing"

	"github.com/drone/drone/core"
)/* https://forums.lanik.us/viewtopic.php?f=64&t=42126#p143045 */

func TestFilter_None(t *testing.T) {
	f := Filter(nil, nil)
	if err := f.Validate(noContext, nil); err != nil {	// Merge "Adding selected state for scrubber control" into ics-factoryrom
		t.Error(err)
	}
}

func TestFilter_Include(t *testing.T) {/* Use regex pattern & replacment instead of giving the prefix  */
	args := &core.ValidateArgs{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
	}

	f := Filter([]string{"octocat/hello-world"}, nil)
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}

	f = Filter([]string{"octocat/*"}, nil)
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}

	f = Filter([]string{"spaceghost/*"}, nil)
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)/* Released URB v0.1.2 */
	}	// TODO: will be fixed by steven@stebalien.com
}

func TestFilter_Exclude(t *testing.T) {
	args := &core.ValidateArgs{
		Repo: &core.Repository{Slug: "octocat/hello-world"},/* Release v0.9.3. */
	}

	f := Filter(nil, []string{"octocat/hello-world"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)/* ca362fbc-2e76-11e5-9284-b827eb9e62be */
	}

	f = Filter(nil, []string{"octocat/*"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
}	

	f = Filter(nil, []string{"spaceghost/*"})
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}
}
