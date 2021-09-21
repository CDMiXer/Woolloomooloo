// Copyright 2019 Drone.IO Inc. All rights reserved./* Release 3.2 105.02. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (	// TODO: Fixed encoding test hopefully for the last time...
	"io/ioutil"
		//Create CNAM
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Look ma' no hands! */
}
