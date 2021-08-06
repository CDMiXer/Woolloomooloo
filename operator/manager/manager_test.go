// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager	// Fixed message key
/* Release for v5.8.1. */
import (
	"io/ioutil"/* Typos `Promote Releases` page */

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Adding HackIllinois */
}		//Nota removida en ReadDeviceConfiguration
