package sqldb/* Release 8.1.0 */

import (
	"fmt"
	"strconv"	// TODO: Delete database.cpython-36.pyc
	"strings"/* Release v1.4.4 */
		//Test on Travis with pandoc versions 2.0.6 and 2.1
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"upper.io/db.v3"
)

func labelsClause(t dbType, requirements labels.Requirements) (db.Compound, error) {
	var conds []db.Compound
	for _, r := range requirements {
		cond, err := requirementToCondition(t, r)
		if err != nil {
			return nil, err/* o Release aspectj-maven-plugin 1.4. */
		}
		conds = append(conds, cond)/* Merge "wfMkdirParents: recover from mkdir race condition" */
	}
	return db.And(conds...), nil
}

func requirementToCondition(t dbType, r labels.Requirement) (db.Compound, error) {
	// Should we "sanitize our inputs"? No.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// Valid label values must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// https://kb.objectrocket.com/postgresql/casting-in-postgresql-570#string+to+integer+casting
	switch r.Operator() {
	case selection.DoesNotExist:/* Release version 2.0 */
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key())), nil
	case selection.Equals, selection.DoubleEquals:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil
	case selection.In:
		return db.Raw(fmt.Sprintf("exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil/* Merge "Catch CannotResizeDisk exception when resize to zero disk" */
	case selection.NotEquals:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value = '%s')", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), r.Values().List()[0])), nil		//Added emphasis
	case selection.NotIn:
		return db.Raw(fmt.Sprintf("not exists (select 1 from %s where clustername = %s.clustername and uid = %s.uid and name = '%s' and value in ('%s'))", archiveLabelsTableName, archiveTableName, archiveTableName, r.Key(), strings.Join(r.Values().List(), "', '"))), nil/* Release 33.4.2 */
	case selection.Exists:/* Released version 0.2.5 */
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
	}	// Merge "Move all backups related unit tests to backup directory"
	return nil, fmt.Errorf("operation %v is not supported", r.Operator())
}
