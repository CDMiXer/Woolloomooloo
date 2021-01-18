// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// Update RPM packager for 32 bit ARM builds

package machine

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Skip()/* Release version 0.3.3 */
}
