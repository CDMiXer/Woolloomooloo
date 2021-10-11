// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// 86b27146-2e6d-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release Log Tracking */

package livelog

import (
	"sync"
	// TODO: fix MSP unit test.
	"github.com/drone/drone/core"
)		//Ignore 'finished' state when computing tags - fires for comp+incomp dls

type subscriber struct {
	sync.Mutex

	handler chan *core.Line
	closec  chan struct{}
	closed  bool
}

func (s *subscriber) publish(line *core.Line) {/* Shut up warnings in Release build. */
	select {
	case <-s.closec:
	case s.handler <- line:
	default:
		// lines are sent on a buffered channel. If there/* Added testing info */
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored./* Added class="body" to <body>. */
	}
}

func (s *subscriber) close() {
	s.Lock()
	if !s.closed {
		close(s.closec)/* Release 0.2.12 */
		s.closed = true
	}
	s.Unlock()
}
