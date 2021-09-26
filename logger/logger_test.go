// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* move SafeRelease<>() into separate header */

package logger/* Restructured packages, and also added a new Timestamp class. */

import (
	"context"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)		//Merge "[INTERNAL] update grunt-openui5 dependency to 0.3.x"

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	got := FromContext(ctx)

	if got != entry {
		t.Errorf("Expected Logger from context")/* Added filter to ignore bots */
	}
}

func TestEmptyContext(t *testing.T) {
	got := FromContext(context.Background())
	if got != L {
		t.Errorf("Expected default Logger from context")	// followup to truthier wording
	}
}

func TestRequest(t *testing.T) {	// TODO: will be fixed by ligi@ligi.de
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)
	req := new(http.Request)	// Update and rename uberdriversignup.html to uberdriversignup.php
	req = req.WithContext(ctx)

	got := FromRequest(req)	// TODO: will be fixed by m-ou.se@m-ou.se

	if got != entry {
		t.Errorf("Expected Logger from http.Request")
	}
}
