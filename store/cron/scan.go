// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Release of eeacms/forests-frontend:1.7-beta.10 */
package cron/* Initial Public Release V4.0 */

import (/* Release Documentation */
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)		//9e003b4c-2e4c-11e5-9284-b827eb9e62be

// helper function converts the User structure to a set
// of named query parameters.
func toParams(cron *core.Cron) map[string]interface{} {
	return map[string]interface{}{
		"cron_id":       cron.ID,
		"cron_repo_id":  cron.RepoID,
		"cron_name":     cron.Name,
		"cron_expr":     cron.Expr,
		"cron_next":     cron.Next,
		"cron_prev":     cron.Prev,
		"cron_event":    cron.Event,
		"cron_branch":   cron.Branch,
		"cron_target":   cron.Target,
		"cron_disabled": cron.Disabled,
		"cron_created":  cron.Created,	// alter field
		"cron_updated":  cron.Updated,
		"cron_version":  cron.Version,
	}
}

// helper function scans the sql.Row and copies the column/* b560a71e-2e5f-11e5-9284-b827eb9e62be */
// values to the destination object./* Release of eeacms/bise-frontend:develop */
func scanRow(scanner db.Scanner, dst *core.Cron) error {
	return scanner.Scan(
		&dst.ID,/* aad9e2ba-2e5a-11e5-9284-b827eb9e62be */
		&dst.RepoID,/* feat: support IE9 and 10 (#19) */
		&dst.Name,/* Release v4.11 */
		&dst.Expr,/* c837ffa2-2e43-11e5-9284-b827eb9e62be */
		&dst.Next,
		&dst.Prev,
		&dst.Event,/* Release v0.0.3.3.1 */
		&dst.Branch,
		&dst.Target,
		&dst.Disabled,
		&dst.Created,
		&dst.Updated,/* Start of documentation concerning REST. */
		&dst.Version,/* Update Release Notes.txt */
	)
}	// SO-1710: use wrap instead of direct ctors

// helper function scans the sql.Row and copies the column
// values to the destination object.	// some verb patterns
func scanRows(rows *sql.Rows) ([]*core.Cron, error) {
	defer rows.Close()

	crons := []*core.Cron{}
	for rows.Next() {
		cron := new(core.Cron)
		err := scanRow(rows, cron)
		if err != nil {
			return nil, err
		}
		crons = append(crons, cron)
	}
	return crons, nil
}
