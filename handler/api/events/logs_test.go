// Copyright 2019 Drone.IO Inc. All rights reserved./* Add L and R key binds to Q and W */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"		//bug(#62):Errores en el panel de control de los centros
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
