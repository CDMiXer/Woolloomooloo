// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release SIIE 3.2 100.02. */
// that can be found in the LICENSE file.
	// TODO: will be fixed by lexy8russo@outlook.com
// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)		//Doc on how to produce compile POM dependencies.
		//Merge branch 'master' into new_bundler
func init() {
	logrus.SetOutput(ioutil.Discard)
}
