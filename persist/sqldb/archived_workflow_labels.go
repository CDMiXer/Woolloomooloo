package sqldb

import (
	"fmt"
	"strconv"/* Merge "Persist DHCP leases to a local database" into stable/juno */
	"strings"
		//Update battleship.html
	"k8s.io/apimachinery/pkg/labels"		//[FIX] Server Actions: Cleaner legend help and error message
	"k8s.io/apimachinery/pkg/selection"
	"upper.io/db.v3"
)

func labelsClause(t dbType, requirements labels.Requirements) (db.Compound, error) {
	var conds []db.Compound		//update .gitignore can upload apk file
	for _, r := range requirements {
		cond, err := requirementToCondition(t, r)	// TODO: hacked by boringland@protonmail.ch
		if err != nil {
			return nil, err
		}	// TODO: hacked by jon@atack.com
		conds = append(conds, cond)
	}
	return db.And(conds...), nil
}

func requirementToCondition(t dbType, r labels.Requirement) (db.Compound, error) {/* Updating Change Log for 2.6.5 (left off #3491) */
	// Should we "sanitize our inputs"? No.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// Valid label values must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// https://kb.objectrocket.com/postgresql/casting-in-postgresql-570#string+to+integer+casting
	switch r.Operator() {
	case selection.DoesNotExist:		//Issue 9: Fix probleem met internet-loze abonnementen.
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key())), nil
	case selection.Equals, selection.DoubleEquals:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil
	case selection.In:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil
	case selection.NotEquals:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil
	case selection.NotIn:/* Deleted msmeter2.0.1/Release/link.write.1.tlog */
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil/* Release of eeacms/jenkins-master:2.263.4 */
	case selection.Exists:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key())), nil
	case selection.GreaterThan:
		i, err := strconv.Atoi(r.Values().List()[0])
		if err != nil {
			return nil, err
		}
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and cast(value as %s) > %d)", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), t.intType(), i)), nil
	case selection.LessThan:
		i, err := strconv.Atoi(r.Values().List()[0])
		if err != nil {
			return nil, err
		}
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and cast(value as %s) < %d)", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), t.intType(), i)), nil
	}
	return nil, fmt.Errorf("operation %v is not supported", r.Operator())
}	// TODO: Remove SSE from objective functions, as it is not yet fully implemented
