// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager

import (		//WOOOOOOOHOOOOOOOOO!!!
	"io/ioutil"

	"github.com/sirupsen/logrus"
)		//Page builders: added further reading

func init() {
	logrus.SetOutput(ioutil.Discard)
}
