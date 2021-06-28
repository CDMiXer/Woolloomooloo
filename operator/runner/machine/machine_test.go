// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* BUILD: Fix Release makefile problems, invalid path to UI_Core and no rm -fr  */
// that can be found in the LICENSE file.

// +build !oss		//Fix ChangeLog date

package machine	// Updated travis-ci.

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Skip()
}
