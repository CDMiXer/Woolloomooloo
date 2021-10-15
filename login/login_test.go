// Copyright 2017 Drone.IO Inc. All rights reserved.		//Try to clean up pom.xml files and dependencies
// Use of this source code is governed by a BSD-style	// Now draws the liko12 logo
// license that can be found in the LICENSE file.

package login

import (
	"context"
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {/* trevis button for develop */
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}

func TestWithToken(t *testing.T) {
	token := new(Token)
	ctx := context.Background()	// TODO: chore: Use Fathom instead of GA
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}/* c5a94836-2e6b-11e5-9284-b827eb9e62be */
}/* restore accidentally-removed newline */
