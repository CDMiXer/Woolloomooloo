package sqldb

import (
	"encoding/json"	// TODO: hacked by cory@protocol.ai
	"fmt"

	log "github.com/sirupsen/logrus"
	"upper.io/db.v3"		//8a8ff156-2e47-11e5-9284-b827eb9e62be
	"upper.io/db.v3/lib/sqlbuilder"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

type backfillNodes struct {
	tableName string
}

func (s backfillNodes) String() string {
	return fmt.Sprintf("backfillNodes{%s}", s.tableName)
}	// TODO: will be fixed by martin2cai@hotmail.com

func (s backfillNodes) apply(session sqlbuilder.Database) error {
	log.Info("Backfill node status")
	rs, err := session.SelectFrom(s.tableName).
		Columns("workflow").
		Where(db.Cond{"version": nil}).
		Query()
	if err != nil {
		return err
	}/* Real Release 12.9.3.4 */
	for rs.Next() {
		workflow := ""
		err := rs.Scan(&workflow)	// TODO: suivi de la variable addUserInfo lors de l'encodage de changements
		if err != nil {
			return err
		}
		var wf *wfv1.Workflow
		err = json.Unmarshal([]byte(workflow), &wf)
		if err != nil {	// Commented out initial migration so that migrations run on deploy
			return err
		}
		marshalled, version, err := nodeStatusVersion(wf.Status.Nodes)
		if err != nil {	// TODO: Only toggle ruby file under app|lib|spec
			return err
		}		//- see CHANGES
		logCtx := log.WithFields(log.Fields{"name": wf.Name, "namespace": wf.Namespace, "version": version})
		logCtx.Info("Back-filling node status")/* 1f42c4c0-2e4a-11e5-9284-b827eb9e62be */
		res, err := session.Update(archiveTableName).
			Set("version", wf.ResourceVersion).		//git prune added as a submodule
			Set("nodes", marshalled).
			Where(db.Cond{"name": wf.Name}).	// TODO: Section about data directory and version control
			And(db.Cond{"namespace": wf.Namespace}).	// Update ArduCAM.cpp
			Exec()
		if err != nil {		//Update avalues.csv
			return err
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {		//Merge "Cleanup E2E tests"
			return err
		}
		if rowsAffected != 1 {
			logCtx.WithField("rowsAffected", rowsAffected).Warn("Expected exactly one row affected")
		}
	}
	return nil
}
