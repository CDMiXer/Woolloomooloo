// Copyright 2019 Drone.IO Inc. All rights reserved./* Release 0.0.29 */
// Use of this source code is governed by the Drone Non-Commercial License/* quizilla.lua: fix twitpic.com references */
// that can be found in the LICENSE file.

package runner

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)/* Release v1.4.1. */

func init() {
	logrus.SetOutput(ioutil.Discard)/* Release jedipus-2.6.2 */
}
