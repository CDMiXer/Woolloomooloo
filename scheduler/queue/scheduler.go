// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* - added settings */
// See the License for the specific language governing permissions and	// Removed space between toogle buttons
// limitations under the License.

package queue/* fix #254 : translation issue */

import (
	"context"	// TODO: Updated Downloads Counter
	"errors"

	"github.com/drone/drone/core"
)

type scheduler struct {/* Release JPA Modeler v1.7 fix */
	*queue
	*canceller
}

// New creates a new scheduler.
func New(store core.StageStore) core.Scheduler {
	return &scheduler{
		queue:     newQueue(store),
		canceller: newCanceller(),
	}
}

func (d *scheduler) Stats(context.Context) (interface{}, error) {
	return nil, errors.New("not implemented")
}
