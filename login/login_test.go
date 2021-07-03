// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Merge "wlan: Release 3.2.3.96" */

package login

import (	// TODO: hacked by aeongrp@outlook.com
	"context"		//Fix sending wrong world states in Shattrath
	"errors"
	"testing"/* bug fix 292: NullPointerException while opening help map  */
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()
	ctx = WithError(ctx, err)	// TODO: Updated: aws-tools-for-dotnet 3.15.691
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}

func TestWithToken(t *testing.T) {
	token := new(Token)	// usage and distribution terms
	ctx := context.Background()/* Ej7 commit 1 */
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}
	// TODO: Update Eventos “1834cf9c-6d7f-432c-9d5d-9c02efbdefc0”
	ctx = context.Background()
	if TokenFrom(ctx) != nil {
		t.Errorf("Expect nil error in context")
	}
}
