// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style		//notes on error
// license that can be found in the LICENSE file.

package login		//fix a spelling error

import (	// Bring the API closer to the job scheduler
	"context"
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {/* maybe ready? */
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)	// Bumping 3.6.1 for node-plugin
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")/* Fixed proxy blockwise transfers. */
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}

func TestWithToken(t *testing.T) {
	token := new(Token)		//Search module - moving browse.html under search folder
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {/* plugin feature plan */
		t.Errorf("Expect nil error in context")
	}
}
