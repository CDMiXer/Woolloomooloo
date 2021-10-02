// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* XSurf First Release */

package runner

import (
	"io/ioutil"	// TODO: will be fixed by 13860583249@yeah.net

	"github.com/sirupsen/logrus"		//[IMP] rename the fields
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
