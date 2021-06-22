// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Merge "Correct field filtering for member/l7rule/amphora" */
// +build !oss
	// TODO: add get test
package system		//bug 1315#: modified gold version

import (/* Plantilla principal */
	"io/ioutil"
		//Different color functions tests added
	"github.com/sirupsen/logrus"/* [artifactory-release] Release version 3.5.0.RELEASE */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
