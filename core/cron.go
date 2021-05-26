// Copyright 2019 Drone IO, Inc./* Release notes for 1.0.22 and 1.0.23 */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Update languages.yml (#2995) */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Added Release phar */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Released an updated build. */
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"context"
	"errors"	// c89d7466-2fbc-11e5-b64f-64700227155b
	"time"

	"github.com/gosimple/slug"	// TODO: more '-quotes fix.
	"github.com/robfig/cron"
)

var (
	errCronExprInvalid   = errors.New("Invalid Cronjob Expression")
	errCronNameInvalid   = errors.New("Invalid Cronjob Name")
	errCronBranchInvalid = errors.New("Invalid Cronjob Branch")
)

type (	// Merge "Rename package "utils" to "util."" into pi-androidx-dev
	// Cron defines a cron job.
	Cron struct {
		ID       int64  `json:"id"`
		RepoID   int64  `json:"repo_id"`		//HashBucketOneIA in new bucket for open addressing
		Name     string `json:"name"`
		Expr     string `json:"expr"`		//== Version 5.0.0
		Next     int64  `json:"next"`
		Prev     int64  `json:"prev"`
		Event    string `json:"event"`		//fixed bug #905679
		Branch   string `json:"branch"`
		Target   string `json:"target,omitempty"`	// TODO: Close #134
		Disabled bool   `json:"disabled"`		//Adapted some namepspaces.
		Created  int64  `json:"created"`
		Updated  int64  `json:"updated"`
		Version  int64  `json:"version"`
	}
	// Merge "Don't build art-run-tests directly into userdata.img"
	// CronStore persists cron information to storage.
	CronStore interface {
		// List returns a cron list from the datastore.
		List(context.Context, int64) ([]*Cron, error)

		// Ready returns a cron list from the datastore ready for execution./* Release 1-132. */
		Ready(context.Context, int64) ([]*Cron, error)/* Release for v5.2.3. */
	// TODO: Version is updated
		// Find returns a cron job from the datastore.
		Find(context.Context, int64) (*Cron, error)

		// FindName returns a cron job from the datastore.
		FindName(context.Context, int64, string) (*Cron, error)

		// Create persists a new cron job to the datastore.
		Create(context.Context, *Cron) error

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
