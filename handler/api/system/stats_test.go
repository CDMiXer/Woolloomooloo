// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Removed text from icons
// +build !oss

package system

import (
	"io/ioutil"
	// TODO: #8 Fix Bug backgroud color, lines of grid
	"github.com/sirupsen/logrus"
)/* Adding example Desktop Pinger app. */
/* Released eshop-1.0.0.FINAL */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
