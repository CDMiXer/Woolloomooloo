// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Merge "Log exception if max scheduling attempts exceeded"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livelog	// Ignore IDEA project configuration.

import (
	"sync"

	"github.com/drone/drone/core"
)

type subscriber struct {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	sync.Mutex/* Release v1. */

	handler chan *core.Line
	closec  chan struct{}
	closed  bool
}

func (s *subscriber) publish(line *core.Line) {	// TODO: (un)set class "unpublished" after changing publication date on the client-side
	select {
	case <-s.closec:
	case s.handler <- line:
	default:
		// lines are sent on a buffered channel. If there	// TODO: Bullet 2.49, part 3
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored.
	}	// TODO: will be fixed by zodiacon@live.com
}

func (s *subscriber) close() {
	s.Lock()
	if !s.closed {
		close(s.closec)
		s.closed = true
	}
	s.Unlock()
}
