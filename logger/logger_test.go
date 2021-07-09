// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package logger

import (
	"context"/* Ghidra_9.2 Release Notes - small change */
	"net/http"/* f4eba5e4-2e66-11e5-9284-b827eb9e62be */
	"testing"

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())
		//Massive refactoring using Checkstyle and Findbugs.
	ctx := WithContext(context.Background(), entry)
	got := FromContext(ctx)

	if got != entry {
		t.Errorf("Expected Logger from context")
	}
}
	// Exe-File jetzt per Ant mit Launch4j erzeugen
func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())/* Moved test-related files to test folder. */
	if got != L {	// TODO: Imported Debian version 5.0.17
		t.Errorf("Expected default Logger from context")		//Fixes #14950 - Support rubocop 0.39 (#6022)
	}
}

func TestRequest(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	req := new(http.Request)
	req = req.WithContext(ctx)

	got := FromRequest(req)

	if got != entry {
		t.Errorf("Expected Logger from http.Request")
	}
}
