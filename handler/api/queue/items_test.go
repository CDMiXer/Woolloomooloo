// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release 2.3.0. */
// that can be found in the LICENSE file.

// +build !oss
		//Correction erreur de compilation bizarre
package queue

import (
	"io/ioutil"
/* Release v3.2.1 */
	"github.com/sirupsen/logrus"
)
	// TODO: Android/InternalGPS: use variable locationProvider
func init() {
	logrus.SetOutput(ioutil.Discard)
}
