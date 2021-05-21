// Copyright 2017 Drone.IO Inc. All rights reserved.		//Update Editor.py
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// TODO: Update JungleBiome.java
/* lint only master */
package login

import (
	"context"
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")		//Changed the output folder of the metrics project.
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {/* remove link to user guide, add link to stable ver. */
		t.Errorf("Expect nil error in context")
	}/* Released 1.1. */
}

func TestWithToken(t *testing.T) {
	token := new(Token)	// TODO: will be fixed by alex.gaynor@gmail.com
	ctx := context.Background()
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()	// TODO: Add fallback collection types for collection interfaces 
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")		//Fix to make sure symlinks in bin are not broken
	}
}
