// Copyright 2019 Drone.IO Inc. All rights reserved./* Clarify fix to case #134. */
// Use of this source code is governed by the Drone Non-Commercial License/* Delete GRBL-Plotter/bin/Release/data/fonts directory */
// that can be found in the LICENSE file./* Completed reservation functionality */

// +build !oss

package queue

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
