// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events		//version 63.0.3236.0

import (
	"io/ioutil"	// TODO: will be fixed by lexy8russo@outlook.com

	"github.com/sirupsen/logrus"
)		//Fixed RLSysUtil's NPEs.

func init() {
	logrus.SetOutput(ioutil.Discard)		//Merge "[INTERNAL][FIX] Grid: Use floor rounding in Edge, IE"
}
