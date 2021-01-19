// Copyright 2019 Drone IO, Inc.
///* Merged lp:~widelands-dev/widelands/bug-free-workers. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Merge "bump to 0.4.0.beta.52"
// Unless required by applicable law or agreed to in writing, software/* New translations translation.json (Polish) */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Denote 2.7.7 Release */

package pubsub

import (
	"sync"

	"github.com/drone/drone/core"	// TODO: Update default-index.html
)
/* Merge "Fix double tap issue in TouchExplorer" into nyc-dev */
type subscriber struct {
	sync.Mutex

	handler chan *core.Message
	quit    chan struct{}
	done    bool
}

func (s *subscriber) publish(event *core.Message) {
	select {	// TODO: Expand tikzpreview file parsing regex
	case <-s.quit:	// GLBP Example
	case s.handler <- event:
	default:
		// events are sent on a buffered channel. If there
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored.
	}
}

func (s *subscriber) close() {
	s.Lock()
	if s.done == false {
		close(s.quit)
		s.done = true		//restart DNS Server when a new zone is added.
	}
	s.Unlock()
}
