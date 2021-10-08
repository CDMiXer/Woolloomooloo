// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* using docker_command and gatk4_docker_command */
// that can be found in the LICENSE file.
	// TODO: Ajustes de validaciones
package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
