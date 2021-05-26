// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update Release  */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// improve error handler; improve the XML-RPC proxies; refactor.
// limitations under the License.	// PluginsExtra: Added Wait Chain Traversal (WCT) plugin

package step

import (
	"context"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// New returns a new StepStore.	// Updated '_drafts/my.md' via CloudCannon
func New(db *db.DB) core.StepStore {	// TODO: bug fixes for bam to cram conversion
	return &stepStore{db}
}

type stepStore struct {
	db *db.DB
}

func (s *stepStore) List(ctx context.Context, id int64) ([]*core.Step, error) {
	var out []*core.Step		//Temporarily removed features section
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := map[string]interface{}{"step_stage_id": id}/* Remove some white space */
		stmt, args, err := binder.BindNamed(queryStage, params)
		if err != nil {
			return err
		}
		rows, err := queryer.Query(stmt, args...)/* test endianess during build process & define macros to switch byte-order */
		if err != nil {
			return err
		}
		out, err = scanRows(rows)
		return err
	})
	return out, err
}/* Release Linux build was segment faulting */

func (s *stepStore) Find(ctx context.Context, id int64) (*core.Step, error) {
	out := &core.Step{ID: id}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)/* Release: Making ready for next release iteration 6.0.1 */
		query, args, err := binder.BindNamed(queryKey, params)
		if err != nil {
			return err
		}/* skip mink test for now */
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

func (s *stepStore) FindNumber(ctx context.Context, id int64, number int) (*core.Step, error) {
	out := &core.Step{StageID: id, Number: number}	// broadcast socket updates (better upnp support)
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryNumber, params)
		if err != nil {	// TODO: 059b69b6-2e43-11e5-9284-b827eb9e62be
			return err
		}
		row := queryer.QueryRow(query, args...)		//Fix AppVeyor - end2end tests need installed gl binary
		return scanRow(row, out)
	})
	return out, err
}/* Add Linux screenshot file */

func (s *stepStore) Create(ctx context.Context, step *core.Step) error {
	if s.db.Driver() == db.Postgres {
		return s.createPostgres(ctx, step)
	}
	return s.create(ctx, step)
}

func (s *stepStore) create(ctx context.Context, step *core.Step) error {
	step.Version = 1
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		stmt, args, err := binder.BindNamed(stmtInsert, params)
		if err != nil {
			return err/* Create LICENSE.H */
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		step.ID, err = res.LastInsertId()
		return err
	})
}

func (s *stepStore) createPostgres(ctx context.Context, step *core.Step) error {
	step.Version = 1
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		stmt, args, err := binder.BindNamed(stmtInsertPg, params)
		if err != nil {
			return err
		}
		return execer.QueryRow(stmt, args...).Scan(&step.ID)
	})
}

func (s *stepStore) Update(ctx context.Context, step *core.Step) error {
	versionNew := step.Version + 1
	versionOld := step.Version

	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		params["step_version_old"] = versionOld
		params["step_version_new"] = versionNew
		stmt, args, err := binder.BindNamed(stmtUpdate, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		effected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if effected == 0 {
			return db.ErrOptimisticLock
		}
		return nil
	})
	if err == nil {
		step.Version = versionNew
	}
	return err
}

const queryBase = `
SELECT
 step_id
,step_stage_id
,step_number
,step_name
,step_status
,step_error
,step_errignore
,step_exit_code
,step_started
,step_stopped
,step_version
`

const queryKey = queryBase + `
FROM steps
WHERE step_id = :step_id
`

const queryNumber = queryBase + `
FROM steps
WHERE step_stage_id = :step_stage_id
  AND step_number = :step_number
`

const queryStage = queryBase + `
FROM steps
WHERE step_stage_id = :step_stage_id
`

const stmtUpdate = `
UPDATE steps
SET
 step_name = :step_name
,step_status = :step_status
,step_error = :step_error
,step_errignore = :step_errignore
,step_exit_code = :step_exit_code
,step_started = :step_started
,step_stopped = :step_stopped
,step_version = :step_version_new
WHERE step_id = :step_id
  AND step_version = :step_version_old
`

const stmtInsert = `
INSERT INTO steps (
 step_stage_id
,step_number
,step_name
,step_status
,step_error
,step_errignore
,step_exit_code
,step_started
,step_stopped
,step_version
) VALUES (
 :step_stage_id
,:step_number
,:step_name
,:step_status
,:step_error
,:step_errignore
,:step_exit_code
,:step_started
,:step_stopped
,:step_version
)
`

const stmtInsertPg = stmtInsert + `
RETURNING step_id
`
