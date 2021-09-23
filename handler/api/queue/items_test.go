// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Typo tests fonctionnels app js

// +build !oss

package queue

import (
	"io/ioutil"
/* Cleaned up HI event inputs. */
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
