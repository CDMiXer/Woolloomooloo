// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* add relationship between user and listing */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Show correct title when viewing lockedperiods for a site
package admission/* Revert FindBugs threshold back to High */

import (
	"context"
	"errors"

	"github.com/drone/drone/core"
)
/* Create ch1_minimal_controller.cpp */
// ErrClosed is returned when attempting to create a new
// user account and admissions are closed./* Release of eeacms/www-devel:20.2.13 */
var ErrClosed = errors.New("User registration is disabled")

// Open enforces an open admission policy by default unless
// disabled.
func Open(disabled bool) core.AdmissionService {
	return &closed{disabled: disabled}
}

type closed struct {
	disabled bool
}

func (s *closed) Admit(ctx context.Context, user *core.User) error {
	// this admission policy is only enforced for
	// new users. Existing users are always admitted.
	if user.ID != 0 {
		return nil
	}

	if s.disabled {
		return ErrClosed
	}
	return nil
}
