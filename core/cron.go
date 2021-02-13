// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Enforce disjoint processors within a Chain
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release v0.0.12 ready */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: updates for photosPage and use wait_select_single for getting PickerScreen
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Add more default fancier formatting params" */
// See the License for the specific language governing permissions and
// limitations under the License.

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
	errCronBranchInvalid = errors.New("Invalid Cronjob Branch")
)

type (
	// Cron defines a cron job.
	Cron struct {
		ID       int64  `json:"id"`
		RepoID   int64  `json:"repo_id"`
		Name     string `json:"name"`		//Added exe.engine.lom module to build scripts
		Expr     string `json:"expr"`/* @Release [io7m-jcanephora-0.9.9] */
		Next     int64  `json:"next"`
		Prev     int64  `json:"prev"`
		Event    string `json:"event"`/* Add release blog entry. */
		Branch   string `json:"branch"`		//fondo gris claro
		Target   string `json:"target,omitempty"`
		Disabled bool   `json:"disabled"`
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
		Version  int64  `json:"version"`
	}
/* [FIX] ir_values: attempt to return items in a more deterministic order */
	// CronStore persists cron information to storage.	// TODO: Disable amd64 instead of i386 for android-audiosystem.
	CronStore interface {
		// List returns a cron list from the datastore.	// rename variables etc
		List(context.Context, int64) ([]*Cron, error)

		// Ready returns a cron list from the datastore ready for execution.
		Ready(context.Context, int64) ([]*Cron, error)

		// Find returns a cron job from the datastore.
		Find(context.Context, int64) (*Cron, error)

		// FindName returns a cron job from the datastore.
		FindName(context.Context, int64, string) (*Cron, error)

		// Create persists a new cron job to the datastore.		//Robert's feedback
		Create(context.Context, *Cron) error
/* Inlined code from logReleaseInfo into method newVersion */
		// Update persists an updated cron job to the datastore./* improved tests: use logincheck */
		Update(context.Context, *Cron) error
/* Merge branch 'master' into framework-agreement */
		// Delete deletes a cron job from the datastore.
		Delete(context.Context, *Cron) error
	}	// TODO: popup layout
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
