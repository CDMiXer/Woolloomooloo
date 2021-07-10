// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {/* Updating build-info/dotnet/roslyn/dev15.6p4 for beta3-62516-04 */
	logrus.SetOutput(ioutil.Discard)
}
