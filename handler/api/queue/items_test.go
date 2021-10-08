// Copyright 2019 Drone.IO Inc. All rights reserved./* Iterators are optimized */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue/* Fast production monitoring of JVM with BPF tools */

import (
	"io/ioutil"
		//Update shifter.v
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* hibernate and DAO is ok */
}
