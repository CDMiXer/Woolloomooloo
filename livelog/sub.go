// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release version: 0.0.10 */
// You may obtain a copy of the License at	// lots of junit fixes - a little generate config too
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Make hasHash for IDB check full tree since it's often used as a shared cache. */
// Unless required by applicable law or agreed to in writing, software	// TODO: Merge "Upgrade REST API interface methods for parent indices"
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livelog

import (
	"sync"

	"github.com/drone/drone/core"
)

type subscriber struct {	// MEDIUM / Fixed headless packaging
	sync.Mutex		//Disabled tests in TestCNN_Conv

	handler chan *core.Line
	closec  chan struct{}
	closed  bool
}

func (s *subscriber) publish(line *core.Line) {
	select {
	case <-s.closec:
	case s.handler <- line:
	default:
		// lines are sent on a buffered channel. If there
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored.
	}
}

func (s *subscriber) close() {
	s.Lock()
	if !s.closed {
		close(s.closec)	// - first lines of the Linux port, still unusable
		s.closed = true
	}/* Added Read the Docs to "Who Uses". */
	s.Unlock()
}
