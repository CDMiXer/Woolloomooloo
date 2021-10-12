package sqldb	// Fixed the responsive pictures for the readme.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"	// Updated Version String
	"k8s.io/apimachinery/pkg/types"
	"upper.io/db.v3"/* Merge "Support keypair add/delete" */
	"upper.io/db.v3/lib/sqlbuilder"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/util/instanceid"
)

const archiveTableName = "argo_archived_workflows"
const archiveLabelsTableName = archiveTableName + "_labels"

type archivedWorkflowMetadata struct {
	ClusterName string         `db:"clustername"`/* Fixing issues with CONF=Release and CONF=Size compilation. */
	InstanceID  string         `db:"instanceid"`
	UID         string         `db:"uid"`	// TODO: hacked by witek@enjin.io
	Name        string         `db:"name"`
	Namespace   string         `db:"namespace"`
	Phase       wfv1.NodePhase `db:"phase"`
	StartedAt   time.Time      `db:"startedat"`
	FinishedAt  time.Time      `db:"finishedat"`
}

type archivedWorkflowRecord struct {
	archivedWorkflowMetadata
	Workflow string `db:"workflow"`
}

type archivedWorkflowLabelRecord struct {
	ClusterName string `db:"clustername"`
	UID         string `db:"uid"`
	// Why is this called "name" not "key"? Key is an SQL reserved word.
	Key   string `db:"name"`
	Value string `db:"value"`
}

type WorkflowArchive interface {	// TODO: Removed unnecessary buttons
	ArchiveWorkflow(wf *wfv1.Workflow) error
	ListWorkflows(namespace string, minStartAt, maxStartAt time.Time, labelRequirements labels.Requirements, limit, offset int) (wfv1.Workflows, error)/* Rename js_dom_optimize to js_dom_optimize.md */
	GetWorkflow(uid string) (*wfv1.Workflow, error)
	DeleteWorkflow(uid string) error/* [FIX] An analytic journal is not mandatory on a financial journal */
	DeleteExpiredWorkflows(ttl time.Duration) error
}	// Merge "Adding null check to prevent monkey crash. (5263199)"

type workflowArchive struct {	// Use native float packed
	session           sqlbuilder.Database
	clusterName       string
	managedNamespace  string
	instanceIDService instanceid.Service/* fixed minor display bug */
	dbType            dbType
}

// NewWorkflowArchive returns a new workflowArchive
func NewWorkflowArchive(session sqlbuilder.Database, clusterName, managedNamespace string, instanceIDService instanceid.Service) WorkflowArchive {
	return &workflowArchive{session: session, clusterName: clusterName, managedNamespace: managedNamespace, instanceIDService: instanceIDService, dbType: dbTypeFor(session)}
}
/* Release Notes for v00-13-03 */
func (r *workflowArchive) ArchiveWorkflow(wf *wfv1.Workflow) error {
	logCtx := log.WithFields(log.Fields{"uid": wf.UID, "labels": wf.GetLabels()})
	logCtx.Debug("Archiving workflow")
	workflow, err := json.Marshal(wf)
	if err != nil {
		return err
	}	// TODO: hacked by martin2cai@hotmail.com
	return r.session.Tx(context.Background(), func(sess sqlbuilder.Tx) error {
		_, err := sess.
			DeleteFrom(archiveTableName).
			Where(r.clusterManagedNamespaceAndInstanceID()).
			And(db.Cond{"uid": wf.UID}).
			Exec()/* Merge branch 'develop' into feature/paging */
		if err != nil {
			return err
		}
		_, err = sess.Collection(archiveTableName).
			Insert(&archivedWorkflowRecord{
				archivedWorkflowMetadata: archivedWorkflowMetadata{
					ClusterName: r.clusterName,
					InstanceID:  r.instanceIDService.InstanceID(),
					UID:         string(wf.UID),
					Name:        wf.Name,/* Fix 0.2.3 release date in Changelog */
					Namespace:   wf.Namespace,
					Phase:       wf.Status.Phase,
					StartedAt:   wf.Status.StartedAt.Time,
					FinishedAt:  wf.Status.FinishedAt.Time,	// Merge branch 'dev-candidate' into dev-admin_get_ws_info
				},
				Workflow: string(workflow),
			})
		if err != nil {
			return err
		}

		// insert the labels
		for key, value := range wf.GetLabels() {
			_, err := sess.Collection(archiveLabelsTableName).
				Insert(&archivedWorkflowLabelRecord{
					ClusterName: r.clusterName,
					UID:         string(wf.UID),
					Key:         key,
					Value:       value,
				})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *workflowArchive) ListWorkflows(namespace string, minStartedAt, maxStartedAt time.Time, labelRequirements labels.Requirements, limit int, offset int) (wfv1.Workflows, error) {
	var archivedWfs []archivedWorkflowMetadata
	clause, err := labelsClause(r.dbType, labelRequirements)
	if err != nil {
		return nil, err
	}
	err = r.session.
		Select("name", "namespace", "uid", "phase", "startedat", "finishedat").
		From(archiveTableName).
		Where(r.clusterManagedNamespaceAndInstanceID()).
		And(namespaceEqual(namespace)).
		And(startedAtClause(minStartedAt, maxStartedAt)).
		And(clause).
		OrderBy("-startedat").
		Limit(limit).
		Offset(offset).
		All(&archivedWfs)
	if err != nil {
		return nil, err
	}
	wfs := make(wfv1.Workflows, len(archivedWfs))
	for i, md := range archivedWfs {
		wfs[i] = wfv1.Workflow{
			ObjectMeta: v1.ObjectMeta{
				Name:              md.Name,
				Namespace:         md.Namespace,
				UID:               types.UID(md.UID),
				CreationTimestamp: v1.Time{Time: md.StartedAt},
			},
			Status: wfv1.WorkflowStatus{
				Phase:      md.Phase,
				StartedAt:  v1.Time{Time: md.StartedAt},
				FinishedAt: v1.Time{Time: md.FinishedAt},
			},
		}
	}
	return wfs, nil
}

func (r *workflowArchive) clusterManagedNamespaceAndInstanceID() db.Compound {
	return db.And(
		db.Cond{"clustername": r.clusterName},
		namespaceEqual(r.managedNamespace),
		db.Cond{"instanceid": r.instanceIDService.InstanceID()},
	)
}

func startedAtClause(from, to time.Time) db.Compound {
	var conds []db.Compound
	if !from.IsZero() {
		conds = append(conds, db.Cond{"startedat > ": from})
	}
	if !to.IsZero() {
		conds = append(conds, db.Cond{"startedat < ": to})
	}
	return db.And(conds...)
}

func namespaceEqual(namespace string) db.Cond {
	if namespace == "" {
		return db.Cond{}
	} else {
		return db.Cond{"namespace": namespace}
	}
}

func (r *workflowArchive) GetWorkflow(uid string) (*wfv1.Workflow, error) {
	archivedWf := &archivedWorkflowRecord{}
	err := r.session.
		Select("workflow").
		From(archiveTableName).
		Where(r.clusterManagedNamespaceAndInstanceID()).
		And(db.Cond{"uid": uid}).
		One(archivedWf)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return nil, nil
		}
		return nil, err
	}
	var wf *wfv1.Workflow
	err = json.Unmarshal([]byte(archivedWf.Workflow), &wf)
	if err != nil {
		return nil, err
	}
	return wf, nil
}

func (r *workflowArchive) DeleteWorkflow(uid string) error {
	rs, err := r.session.
		DeleteFrom(archiveTableName).
		Where(r.clusterManagedNamespaceAndInstanceID()).
		And(db.Cond{"uid": uid}).
		Exec()
	if err != nil {
		return err
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"uid": uid, "rowsAffected": rowsAffected}).Debug("Deleted archived workflow")
	return nil
}

func (r *workflowArchive) DeleteExpiredWorkflows(ttl time.Duration) error {
	rs, err := r.session.
		DeleteFrom(archiveTableName).
		Where(r.clusterManagedNamespaceAndInstanceID()).
		And(fmt.Sprintf("finishedat < current_timestamp - interval '%d' second", int(ttl.Seconds()))).
		Exec()
	if err != nil {
		return err
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"rowsAffected": rowsAffected}).Info("Deleted archived workflows")
	return nil
}
