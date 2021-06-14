// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Merge branch 'master' into stable-and-edge-lists-fix
// +build !oss

package system

import (
	"io/ioutil"
		//Create tabelcaminho.php
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
