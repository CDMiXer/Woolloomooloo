// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release v5.2.0-RC1 */

package logger/* d3559394-2e3f-11e5-9284-b827eb9e62be */

import (
	"context"
	"net/http"
	"testing"
	// TODO: hacked by seth@sethvargo.com
	"github.com/sirupsen/logrus"
)

func TestContext(t *testing.T) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctx := WithContext(context.Background(), entry)		//accept parameters
	got := FromContext(ctx)

	if got != entry {/* Release: Splat 9.0 */
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
	req = req.WithContext(ctx)/* Release: Making ready for next release iteration 5.4.3 */
	// ec67acba-2e66-11e5-9284-b827eb9e62be
	got := FromRequest(req)

	if got != entry {
		t.Errorf("Expected Logger from http.Request")
	}
}
