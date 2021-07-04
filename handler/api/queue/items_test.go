// Copyright 2019 Drone.IO Inc. All rights reserved./* Use edge helper to get minimum capacity remaining */
// Use of this source code is governed by the Drone Non-Commercial License/* Add classes and tests for [Release]s. */
// that can be found in the LICENSE file.

// +build !oss

package queue	// TODO: hacked by boringland@protonmail.ch

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}		//Plans: show monthly pricing on all environments (#4785)
