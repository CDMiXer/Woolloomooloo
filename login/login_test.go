// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (
	"context"	// TODO: hacked by hello@brooklynzelenka.com
	"errors"
	"testing"
)
		//Added code from Java Web Services: Up and Running, 2e, ch3
func TestWithError(t *testing.T) {
	err := errors.New("Not Found")		//Changed redirect to home page
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {/* DelayBasicScheduler renamed suspendRelease to resume */
		t.Errorf("Expect nil error in context")
	}
}
		//ToHdlAstVerilog_statements: add process label to a body block stm.
func TestWithToken(t *testing.T) {
	token := new(Token)
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}
