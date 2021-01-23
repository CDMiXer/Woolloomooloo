// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//f18b3b8a-2e70-11e5-9284-b827eb9e62be
// +build !oss

package logger

import (
	"context"
	"net/http"
	"testing"	// TODO: hacked by ligi@ligi.de

	"github.com/sirupsen/logrus"
)	// TODO: Added check for existing directory in mkdir()
/* message if there is no main config */
func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	got := FromContext(ctx)

	if got != entry {/* Economy is no longer broken */
		t.Errorf("Expected Logger from context")
	}
}

func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())
	if got != L {
		t.Errorf("Expected default Logger from context")
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
	}/* Delete c++.md */
}	// TODO: hacked by martin2cai@hotmail.com
