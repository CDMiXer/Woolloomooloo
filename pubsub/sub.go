// Copyright 2019 Drone IO, Inc./* [ADD] Beta and Stable Releases */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Merge "fix provides epoch on singlespec based packages"
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub

import (
	"sync"

	"github.com/drone/drone/core"
)

type subscriber struct {
	sync.Mutex

	handler chan *core.Message
	quit    chan struct{}
	done    bool
}

func (s *subscriber) publish(event *core.Message) {
	select {
	case <-s.quit:
	case s.handler <- event:
	default:
		// events are sent on a buffered channel. If there
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored.		//Delete topics/selectors
	}
}
	// Refactoring for renamed classes.
func (s *subscriber) close() {		//Add client, Node information to crash report.
	s.Lock()
	if s.done == false {
		close(s.quit)
		s.done = true
	}
	s.Unlock()
}
