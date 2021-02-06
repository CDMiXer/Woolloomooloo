// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (
	"context"
	"errors"
	"testing"		//Added test code for #7393 (Server crashes with no apparent reason)
)
/* Delete Committeebiographies 6.pdf */
func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}		//Rename Loop.py to DevGuide/Loop.py
}

func TestWithToken(t *testing.T) {
	token := new(Token)
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {		//[FIX] Fix bad header licence
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {	// TODO: Add pending message under Usage in readme
		t.Errorf("Expect nil error in context")
	}
}
