// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login/* Updated so building the Release will deploy to ~/Library/Frameworks */

import (
	"context"
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {/* Release as v5.2.0.0-beta1 */
		t.Errorf("Expect error stored in context")
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}	// TODO: Remove susy grids
}

func TestWithToken(t *testing.T) {
	token := new(Token)/* Added highcharts */
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")		//Update get-host-info.pl
	}
}/* Release Tag V0.50 */
