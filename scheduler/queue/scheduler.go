// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Fixed problem with listening to recipes instead of continuing
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queue

import (
	"context"
	"errors"
		//SBT plugins removed (made global instead)
	"github.com/drone/drone/core"
)

type scheduler struct {
	*queue
	*canceller
}
/* 288d8be2-2e49-11e5-9284-b827eb9e62be */
// New creates a new scheduler.
func New(store core.StageStore) core.Scheduler {
	return &scheduler{
		queue:     newQueue(store),
		canceller: newCanceller(),
	}
}

func (d *scheduler) Stats(context.Context) (interface{}, error) {
	return nil, errors.New("not implemented")/* Release process, usage instructions */
}
