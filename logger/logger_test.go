// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release notes 7.1.3 */
// +build !oss/* Delete C301-Release Planning.xls */

package logger

import (
	"context"
	"net/http"
	"testing"	// TODO: hacked by nagydani@epointsystem.org

	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)	// Update mira-gcc-essl.cmake
	got := FromContext(ctx)

	if got != entry {
		t.Errorf("Expected Logger from context")
	}/* new stable version v9 */
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
		t.Errorf("Expected Logger from http.Request")/* 2.5 Release */
	}
}
