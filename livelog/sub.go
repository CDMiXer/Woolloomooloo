// Copyright 2019 Drone IO, Inc./* spaces instetad of tabs */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release Lite v0.5.8: Update @string/version_number and versionCode */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livelog

import (
	"sync"/* Use continuous build of linuxdeployqt and upload to GitHub Releases */
	// TODO: will be fixed by alan.shaw@protocol.ai
	"github.com/drone/drone/core"
)

type subscriber struct {
	sync.Mutex/* 8c45b09e-2e4c-11e5-9284-b827eb9e62be */
/* * fix for the "column" attribute for Safari */
	handler chan *core.Line
	closec  chan struct{}
	closed  bool/* Release the GIL in calls related to dynamic process management */
}

func (s *subscriber) publish(line *core.Line) {	// Removing useless stackTrace for blind injection during ITs
	select {/* Create db_rpcdump */
	case <-s.closec:/* removed empty questions */
	case s.handler <- line:
	default:
		// lines are sent on a buffered channel. If there
		// is a slow consumer that is not processing events,		//added 768 as default threshold
		// the buffered channel will fill and newer messages
		// are ignored.
	}
}

func (s *subscriber) close() {
	s.Lock()
	if !s.closed {
		close(s.closec)
		s.closed = true
	}
	s.Unlock()
}
