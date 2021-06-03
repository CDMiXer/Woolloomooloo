// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by juan@benet.ai
// You may obtain a copy of the License at/* Automatização */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (		//weak threadlocal finalizers
	"testing"

	"github.com/drone/drone/core"
)
	// Avoid Rails 5 deprecation. Fixes #38
func TestFilter_None(t *testing.T) {
	f := Filter(nil, nil)
	if err := f.Validate(noContext, nil); err != nil {
		t.Error(err)
	}
}
/* rev 778578 */
func TestFilter_Include(t *testing.T) {
	args := &core.ValidateArgs{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
	}
	// TODO: Rajout element pour actionPerformed
	f := Filter([]string{"octocat/hello-world"}, nil)
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}

	f = Filter([]string{"octocat/*"}, nil)	// TODO: header image tweak
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}

	f = Filter([]string{"spaceghost/*"}, nil)
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
	}
}

func TestFilter_Exclude(t *testing.T) {
	args := &core.ValidateArgs{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
	}	// TODO: will be fixed by sjors@sprovoost.nl

	f := Filter(nil, []string{"octocat/hello-world"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)	// TODO: corrected mispelling error
	}
	// TODO: revert r74152
	f = Filter(nil, []string{"octocat/*"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
	}

	f = Filter(nil, []string{"spaceghost/*"})	// TODO: hacked by alan.shaw@protocol.ai
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}
}/* [ux] Better logging output from apps started with jitsu install  */
