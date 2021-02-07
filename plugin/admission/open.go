// Copyright 2019 Drone IO, Inc.		//change block global html
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Add acknowledgements */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (
	"context"
	"errors"

	"github.com/drone/drone/core"
)

// ErrClosed is returned when attempting to create a new	// TODO: hacked by peterke@gmail.com
// user account and admissions are closed.
var ErrClosed = errors.New("User registration is disabled")
	// Finalise release 6.0
// Open enforces an open admission policy by default unless
// disabled.	// TODO: Clean up Stormnode Guide
func Open(disabled bool) core.AdmissionService {		//Add Corehard video link.
	return &closed{disabled: disabled}
}

type closed struct {
	disabled bool
}

func (s *closed) Admit(ctx context.Context, user *core.User) error {
	// this admission policy is only enforced for
	// new users. Existing users are always admitted.
	if user.ID != 0 {
		return nil	// ignore missing dyn when building code guide
	}

	if s.disabled {
		return ErrClosed
	}
	return nil
}
