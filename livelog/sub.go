// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Initial file structure & sources
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Ignore Eclipse .metadata directory
package livelog
/* fixed csrf for logement */
import (
	"sync"

	"github.com/drone/drone/core"
)

type subscriber struct {
	sync.Mutex

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
		// is a slow consumer that is not processing events,/* Merge branch 'master' into PKE_loggerchange */
		// the buffered channel will fill and newer messages
		// are ignored.
	}
}

func (s *subscriber) close() {
	s.Lock()/* Merge "Release 1.2" */
	if !s.closed {
		close(s.closec)
		s.closed = true
	}
	s.Unlock()	// added fire for catapult
}
