// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner	// TODO: will be fixed by zhen6939@gmail.com
	// TODO: will be fixed by yuvalalaluf@gmail.com
import (/* Rename README.{js -> md} */
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
