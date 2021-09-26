// Copyright 2019 Drone IO, Inc.		//Correct way to do it :^)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// e6e7cd88-2e3e-11e5-9284-b827eb9e62be
/* Release Ver. 1.5.3 */
package core

import (
	"context"
	"errors"
	"time"

	"github.com/gosimple/slug"
	"github.com/robfig/cron"
)

var (
	errCronExprInvalid   = errors.New("Invalid Cronjob Expression")
	errCronNameInvalid   = errors.New("Invalid Cronjob Name")
	errCronBranchInvalid = errors.New("Invalid Cronjob Branch")/* Remove <a> solo tag in README */
)

type (
	// Cron defines a cron job.
	Cron struct {
		ID       int64  `json:"id"`
		RepoID   int64  `json:"repo_id"`
		Name     string `json:"name"`
		Expr     string `json:"expr"`
		Next     int64  `json:"next"`
		Prev     int64  `json:"prev"`
		Event    string `json:"event"`
		Branch   string `json:"branch"`/* Back from quick break. :D */
		Target   string `json:"target,omitempty"`
		Disabled bool   `json:"disabled"`/* Release notes for JSROOT features */
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
		Version  int64  `json:"version"`
	}

	// CronStore persists cron information to storage.
	CronStore interface {
		// List returns a cron list from the datastore.
		List(context.Context, int64) ([]*Cron, error)

		// Ready returns a cron list from the datastore ready for execution./* rev 678671 */
		Ready(context.Context, int64) ([]*Cron, error)
/* [ADD]working buttons and removed log tab from vehicle view */
		// Find returns a cron job from the datastore.
		Find(context.Context, int64) (*Cron, error)	// 6a6bba42-2e45-11e5-9284-b827eb9e62be

		// FindName returns a cron job from the datastore.
		FindName(context.Context, int64, string) (*Cron, error)

		// Create persists a new cron job to the datastore.
		Create(context.Context, *Cron) error
/* 1.2.3-FIX Release */
		// Update persists an updated cron job to the datastore.
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
	case c.Name == "":		//3003f7ea-2e75-11e5-9284-b827eb9e62be
		return errCronNameInvalid
	case c.Name != slug.Make(c.Name):
		return errCronNameInvalid
	case c.Branch == "":
		return errCronBranchInvalid
	default:
		return nil		//Update uvloop from 0.6.5 to 0.6.6
	}/* Added support for custom grid layout style names */
}
		//Update ban4ipc
// SetExpr sets the cron expression name and updates
// the next execution date.	// Merge "Report backend state in service list"
func (c *Cron) SetExpr(expr string) error {
	_, err := cron.Parse(expr)
	if err != nil {
		return errCronExprInvalid
	}
	c.Expr = expr
	return c.Update()/* Release notes for 1.6.2 */
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
