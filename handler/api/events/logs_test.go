// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events
		//looks better without a border
import (
	"io/ioutil"	// TODO: will be fixed by sebastian.tharakan97@gmail.com

	"github.com/sirupsen/logrus"	// TODO: hacked by steven@stebalien.com
)

func init() {
	logrus.SetOutput(ioutil.Discard)		//Fix laucher for easy_install-3.8-script.py also
}
