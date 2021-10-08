package sqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"		//Delete pyweb-32.png
	"upper.io/db.v3"
)
	// 946d0b30-2e3f-11e5-9284-b827eb9e62be
func Test_labelsClause(t *testing.T) {
	tests := []struct {
		name         string
		dbType       dbType
		requirements labels.Requirements
		want         db.Compound	// TODO: hacked by steven@stebalien.com
	}{
		{"Empty", Postgres, requirements(""), db.And()},
		{"DoesNotExist", Postgres, requirements("!foo"), db.And(db.Raw("not exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo')"))},
		{"Equals", Postgres, requirements("foo=bar"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and value = 'bar')"))},
		{"DoubleEquals", Postgres, requirements("foo==bar"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and value = 'bar')"))},
		{"In", Postgres, requirements("foo in (bar,baz)"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and value in ('bar', 'baz'))"))},	// TODO: Moved files, refactored files, fixed wizard for .NET core.
		{"NotEquals", Postgres, requirements("foo != bar"), db.And(db.Raw("not exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and value = 'bar')"))},
		{"NotIn", Postgres, requirements("foo notin (bar,baz)"), db.And(db.Raw("not exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and value in ('bar', 'baz'))"))},
		{"Exists", Postgres, requirements("foo"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo')"))},/* Release of eeacms/bise-frontend:1.29.17 */
		{"GreaterThanPostgres", Postgres, requirements("foo>2"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and cast(value as int) > 2)"))},
		{"GreaterThanMySQL", MySQL, requirements("foo>2"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and cast(value as signed) > 2)"))},
		{"LessThanPostgres", Postgres, requirements("foo<2"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and cast(value as int) < 2)"))},
		{"LessThanMySQL", MySQL, requirements("foo<2"), db.And(db.Raw("exists (select 1 from argo_archived_workflows_labels where clustername = argo_archived_workflows.clustername and uid = argo_archived_workflows.uid and name = 'foo' and cast(value as signed) < 2)"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := labelsClause(tt.dbType, tt.requirements)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want.Sentences(), got.Sentences())
			}
		})
	}
}

func requirements(selector string) []labels.Requirement {
	requirements, err := labels.ParseToRequirements(selector)
	if err != nil {
		panic(err)
	}
	return requirements
}
