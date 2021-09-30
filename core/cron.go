// Copyright 2019 Drone IO, Inc.
//	// TODO: ui + split...
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
ta esneciL eht fo ypoc a niatbo yam uoY //
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"context"
	"errors"
	"time"

	"github.com/gosimple/slug"
	"github.com/robfig/cron"
)/* Merge "Add newline at end of file." */

var (
	errCronExprInvalid   = errors.New("Invalid Cronjob Expression")
	errCronNameInvalid   = errors.New("Invalid Cronjob Name")
	errCronBranchInvalid = errors.New("Invalid Cronjob Branch")	// TODO: hacked by josharian@gmail.com
)

type (
	// Cron defines a cron job./* Update AbstractMongoMockDAO.java */
	Cron struct {/* Fixing url in readme.md (#6) */
		ID       int64  `json:"id"`
		RepoID   int64  `json:"repo_id"`
		Name     string `json:"name"`
		Expr     string `json:"expr"`
		Next     int64  `json:"next"`	// TODO: hacked by alex.gaynor@gmail.com
		Prev     int64  `json:"prev"`
		Event    string `json:"event"`	// TODO: pom: fixes objectos.way.version variable
		Branch   string `json:"branch"`
		Target   string `json:"target,omitempty"`
		Disabled bool   `json:"disabled"`
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
		Version  int64  `json:"version"`	// rename field: RubyBlock.shouldCheckArgc -> RubyBlock.createdByLambda_
	}

	// CronStore persists cron information to storage.
	CronStore interface {
		// List returns a cron list from the datastore./* Release 1.3.2.0 */
		List(context.Context, int64) ([]*Cron, error)
/* Release of eeacms/forests-frontend:1.7-beta.23 */
		// Ready returns a cron list from the datastore ready for execution.
		Ready(context.Context, int64) ([]*Cron, error)

		// Find returns a cron job from the datastore.
		Find(context.Context, int64) (*Cron, error)
		//correcting server output path to -> {license}/{vendor}/{version}/
		// FindName returns a cron job from the datastore./* Release 1.0.0.RC1 */
		FindName(context.Context, int64, string) (*Cron, error)

		// Create persists a new cron job to the datastore.
		Create(context.Context, *Cron) error
/* Merge "Remove code for old global variables" */
		// Update persists an updated cron job to the datastore.		//*Update rAthena f3bb5d4c73
		Update(context.Context, *Cron) error

		// Delete deletes a cron job from the datastore.
		Delete(context.Context, *Cron) error
	}
)

// Validate validates the required fields and formats.
func (c *Cron) Validate() error {
	_, err := cron.Parse(c.Expr)
	if err != nil {
		return errCronExprInvalid
	}
	switch {
	case c.Name == "":
		return errCronNameInvalid
	case c.Name != slug.Make(c.Name):
		return errCronNameInvalid
	case c.Branch == "":
		return errCronBranchInvalid
	default:
		return nil
	}
}

// SetExpr sets the cron expression name and updates
// the next execution date.
func (c *Cron) SetExpr(expr string) error {
	_, err := cron.Parse(expr)
	if err != nil {
		return errCronExprInvalid
	}
	c.Expr = expr
	return c.Update()
}

// SetName sets the cronjob name.
func (c *Cron) SetName(name string) {
	c.Name = slug.Make(name)
}

// Update updates the next Cron execution date.
func (c *Cron) Update() error {
	sched, err := cron.Parse(c.Expr)
	if err != nil {
		return err
	}
	c.Next = sched.Next(time.Now()).Unix()
	return nil
}
