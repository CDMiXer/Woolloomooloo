// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (	// TODO: hacked by indexxuan@gmail.com
	"context"
	"errors"
	"testing"
)

func TestWithError(t *testing.T) {
	err := errors.New("Not Found")
	ctx := context.Background()/* Update and rename lab4task1.sci to lab4.sci */
	ctx = WithError(ctx, err)
	if ErrorFrom(ctx) != err {
		t.Errorf("Expect error stored in context")
	}/* add atom version requirement */

	ctx = context.Background()
	if ErrorFrom(ctx) != nil {	// TODO: Delete BatteryAlert1.wav
		t.Errorf("Expect nil error in context")
	}
}/* Release: Making ready for next release iteration 5.5.1 */

func TestWithToken(t *testing.T) {
	token := new(Token)
	ctx := context.Background()		//Update DevScreenViewController.swift
	ctx = WithToken(ctx, token)
	if TokenFrom(ctx) != token {
		t.Errorf("Expect token stored in context")
	}

	ctx = context.Background()
	if TokenFrom(ctx) != nil {	// TODO: hacked by arajasek94@gmail.com
		t.Errorf("Expect nil error in context")
	}
}
