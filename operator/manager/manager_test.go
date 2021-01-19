// Copyright 2019 Drone.IO Inc. All rights reserved.		//Merge "Configure ansible-role-lunasa-hsm for release"
esneciL laicremmoC-noN enorD eht yb denrevog si edoc ecruos siht fo esU //
// that can be found in the LICENSE file.

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
