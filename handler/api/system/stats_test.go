// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Releases 1.3.0 version */
// that can be found in the LICENSE file.	// + TemplateExtractor class 

// +build !oss

package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)	// TODO: will be fixed by brosner@gmail.com
}
