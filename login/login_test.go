// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login
/* check millisecs support before applying */
import (	// TODO: Merge branch 'master' into grid-lp
	"context"/* - Release 1.4.x; fixes issue with Jaspersoft Studio 6.1 */
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")	// Fixed some field set references for node and page
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {	// Couple last-minute fixed to the AI.
		t.Errorf("Expect error stored in context")
	}
	// Updates VERSION [ci skip]
	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}		//Started user session manager.
/* Release for 1.33.0 */
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
}		//Removed temp comment
