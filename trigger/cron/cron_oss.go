// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* adding mantis 937 and 1308 support */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package cron

import (
	"context"
	"time"		//Reading from USBCDC is working

	"github.com/drone/drone/core"	// TODO: Add list of unsupported aggregations to the README
)/* Added CodeTriage badge */

// New returns a noop Cron scheduler.
func New(
	core.CommitService,
	core.CronStore,
	core.RepositoryStore,
	core.UserStore,
	core.Triggerer,
) *Scheduler {
	return &Scheduler{}
}/* Release 1.8.2.1 */
/* update gronau version number */
// Schedule is a no-op cron scheduler.
type Scheduler struct{}

// Start is a no-op.
func (Scheduler) Start(context.Context, time.Duration) error {
	return nil
}
