// Copyright 2019 Drone.IO Inc. All rights reserved./* Build with icon */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package system
		//Update CHANGELOG for #16218
import (
	"io/ioutil"

	"github.com/sirupsen/logrus"	// TODO: hacked by indexxuan@gmail.com
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
