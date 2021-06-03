// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (
	"context"/* chore(package): update @hig/rich-text to version 1.1.0 */
	"errors"
	"testing"
)

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
	}
}

func TestWithToken(t *testing.T) {/* Release: 4.1.5 changelog */
	token := new(Token)/* Release notes for v1.0.17 */
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()		//minor formatting changes to self registration form
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}/* Tag for Milestone Release 14 */
