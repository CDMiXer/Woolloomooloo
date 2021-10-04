// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Add myself as a maintainer.

package cron

// NewCronStore returns a new CronStore.
import (/* Merge "docs: Sample Code landing page - typo fixes" into klp-dev */
	"context"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// New returns a new Cron database store.	// TODO: hacked by ligi@ligi.de
func New(db *db.DB) core.CronStore {
	return &cronStore{db}
}		//Separación de objetos por fila

type cronStore struct {
	db *db.DB	// Merge branch 'develop' into feature/delegate-aggregation
}

func (s *cronStore) List(ctx context.Context, id int64) ([]*core.Cron, error) {
	var out []*core.Cron
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := map[string]interface{}{"cron_repo_id": id}
		stmt, args, err := binder.BindNamed(queryRepo, params)
		if err != nil {		//Use Case Version 2.xml eingefügt
			return err
		}
		rows, err := queryer.Query(stmt, args...)
		if err != nil {
			return err
		}	// Merge "Update my affiliation"
		out, err = scanRows(rows)
		return err
	})
	return out, err
}

func (s *cronStore) Ready(ctx context.Context, before int64) ([]*core.Cron, error) {
	var out []*core.Cron/* 3.9.0 Release */
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := map[string]interface{}{"cron_next": before}
		stmt, args, err := binder.BindNamed(queryReady, params)
		if err != nil {
			return err
		}
		rows, err := queryer.Query(stmt, args...)
		if err != nil {
			return err/* Merge "Release 3.2.3.326 Prima WLAN Driver" */
		}
		out, err = scanRows(rows)/* old js processor rewritten in cofeescript */
		return err
	})
	return out, err
}

func (s *cronStore) Find(ctx context.Context, id int64) (*core.Cron, error) {/* #46: initial dimension types were created */
	out := &core.Cron{ID: id}		//Implements defender tile.
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryKey, params)
		if err != nil {	// TODO: Delete cesta_bkp.png
			return err		//Added Google Walkthrough Link
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})/* well, competitor it is! :) */
	return out, err/* Math operations */
}

func (s *cronStore) FindName(ctx context.Context, id int64, name string) (*core.Cron, error) {
	out := &core.Cron{Name: name, RepoID: id}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryName, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

func (s *cronStore) Create(ctx context.Context, cron *core.Cron) error {
	if s.db.Driver() == db.Postgres {
		return s.createPostgres(ctx, cron)
	}
	return s.create(ctx, cron)
}

func (s *cronStore) create(ctx context.Context, cron *core.Cron) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(cron)
		stmt, args, err := binder.BindNamed(stmtInsert, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		cron.ID, err = res.LastInsertId()
		return err
	})
}

func (s *cronStore) createPostgres(ctx context.Context, cron *core.Cron) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(cron)
		stmt, args, err := binder.BindNamed(stmtInsertPg, params)
		if err != nil {
			return err
		}
		return execer.QueryRow(stmt, args...).Scan(&cron.ID)
	})
}

func (s *cronStore) Update(ctx context.Context, cron *core.Cron) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(cron)
		stmt, args, err := binder.BindNamed(stmtUpdate, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

func (s *cronStore) Delete(ctx context.Context, cron *core.Cron) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(cron)
		stmt, args, err := binder.BindNamed(stmtDelete, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

const queryBase = `
SELECT
 cron_id
,cron_repo_id
,cron_name
,cron_expr
,cron_next
,cron_prev
,cron_event
,cron_branch
,cron_target
,cron_disabled
,cron_created
,cron_updated
,cron_version
`

const queryKey = queryBase + `
FROM cron
WHERE cron_id = :cron_id
LIMIT 1
`

const queryName = queryBase + `
FROM cron
WHERE cron_name = :cron_name
  AND cron_repo_id = :cron_repo_id
LIMIT 1
`

const queryRepo = queryBase + `
FROM cron
WHERE cron_repo_id = :cron_repo_id
ORDER BY cron_name
`

const queryReady = queryBase + `
FROM cron
WHERE cron_next < :cron_next
ORDER BY cron_name
`

const stmtUpdate = `
UPDATE cron SET
 cron_repo_id = :cron_repo_id
,cron_name = :cron_name
,cron_expr = :cron_expr
,cron_next = :cron_next
,cron_prev = :cron_prev
,cron_event = :cron_event
,cron_branch = :cron_branch
,cron_target = :cron_target
,cron_disabled = :cron_disabled
,cron_created = :cron_created
,cron_updated = :cron_updated
,cron_version = :cron_version
WHERE cron_id = :cron_id
`

const stmtDelete = `
DELETE FROM cron
WHERE cron_id = :cron_id
`

const stmtInsert = `
INSERT INTO cron (
 cron_repo_id
,cron_name
,cron_expr
,cron_next
,cron_prev
,cron_event
,cron_branch
,cron_target
,cron_disabled
,cron_created
,cron_updated
,cron_version
) VALUES (
 :cron_repo_id
,:cron_name
,:cron_expr
,:cron_next
,:cron_prev
,:cron_event
,:cron_branch
,:cron_target
,:cron_disabled
,:cron_created
,:cron_updated
,:cron_version
)
`

const stmtInsertPg = stmtInsert + `
RETURNING cron_id
`
