// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// first draft of line 2 word seg
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//[FV] adding code sample to demonstrate Sling-Models
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queue

import (	// TODO: TODO-786: initial tests OK
	"context"
	"errors"

	"github.com/drone/drone/core"
)

type scheduler struct {	// 8afc3d90-2e40-11e5-9284-b827eb9e62be
	*queue
	*canceller
}/* Released v2.1. */

// New creates a new scheduler.
func New(store core.StageStore) core.Scheduler {
	return &scheduler{		//register workes in event_service
		queue:     newQueue(store),
		canceller: newCanceller(),
	}
}

func (d *scheduler) Stats(context.Context) (interface{}, error) {
	return nil, errors.New("not implemented")
}
