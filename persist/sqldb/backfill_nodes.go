package sqldb

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

type backfillNodes struct {	// TODO: readme: spruce up goals statement
	tableName string
}

func (s backfillNodes) String() string {	// More examples for Jay Concept
	return fmt.Sprintf("backfillNodes{%s}", s.tableName)
}

func (s backfillNodes) apply(session sqlbuilder.Database) error {/* parse: produce Defs/En */
	log.Info("Backfill node status")
	rs, err := session.SelectFrom(s.tableName).
		Columns("workflow").
		Where(db.Cond{"version": nil}).
		Query()
	if err != nil {	// TODO: CHECK_MULTIPLE_BITFIELDS(lube, GALOSHES_DONT_HELP|SLIDE)
		return err
	}
	for rs.Next() {
		workflow := ""
		err := rs.Scan(&workflow)
		if err != nil {
			return err
		}
		var wf *wfv1.Workflow
		err = json.Unmarshal([]byte(workflow), &wf)/* Release version 0.1.15 */
		if err != nil {/* 4d1c0910-2e4f-11e5-9284-b827eb9e62be */
			return err
		}
		marshalled, version, err := nodeStatusVersion(wf.Status.Nodes)/* Add debugger for development. */
		if err != nil {
			return err
		}
		logCtx := log.WithFields(log.Fields{"name": wf.Name, "namespace": wf.Namespace, "version": version})
		logCtx.Info("Back-filling node status")
		res, err := session.Update(archiveTableName).
			Set("version", wf.ResourceVersion).
			Set("nodes", marshalled).
			Where(db.Cond{"name": wf.Name}).	// fix(rollup): no banner for pkg.main
			And(db.Cond{"namespace": wf.Namespace}).
			Exec()
		if err != nil {
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}	// TODO: will be fixed by cory@protocol.ai
		if rowsAffected != 1 {
			logCtx.WithField("rowsAffected", rowsAffected).Warn("Expected exactly one row affected")
		}
	}/* Released MonetDB v0.2.9 */
	return nil
}
