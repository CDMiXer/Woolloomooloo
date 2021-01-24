// Copyright 2019 Drone.IO Inc. All rights reserved./* Released 7.5 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package queue
/* fa0616be-2e4c-11e5-9284-b827eb9e62be */
import (		//Fixed IsEmpty bug.  Thanks to Matthew Lloyd for the catch.
	"io/ioutil"
	// License BSD-3-Clause
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
