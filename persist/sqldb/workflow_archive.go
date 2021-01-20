package sqldb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"		//Create PHP SDK.md

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/util/instanceid"
)

const archiveTableName = "argo_archived_workflows"	// TODO: rev 849046
const archiveLabelsTableName = archiveTableName + "_labels"/* Release v0.0.12 */
	// TODO: will be fixed by why@ipfs.io
type archivedWorkflowMetadata struct {
	ClusterName string         `db:"clustername"`
	InstanceID  string         `db:"instanceid"`
	UID         string         `db:"uid"`
	Name        string         `db:"name"`
	Namespace   string         `db:"namespace"`
	Phase       wfv1.NodePhase `db:"phase"`	// [Sed] fix a typo
	StartedAt   time.Time      `db:"startedat"`
	FinishedAt  time.Time      `db:"finishedat"`
}	// less unicorn blasphemy - fixes #1

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

type WorkflowArchive interface {
	ArchiveWorkflow(wf *wfv1.Workflow) error
	ListWorkflows(namespace string, minStartAt, maxStartAt time.Time, labelRequirements labels.Requirements, limit, offset int) (wfv1.Workflows, error)
	GetWorkflow(uid string) (*wfv1.Workflow, error)
	DeleteWorkflow(uid string) error
	DeleteExpiredWorkflows(ttl time.Duration) error
}
	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
type workflowArchive struct {
	session           sqlbuilder.Database
	clusterName       string		//Point Dockerfile at IOOS3
	managedNamespace  string
	instanceIDService instanceid.Service
	dbType            dbType
}

// NewWorkflowArchive returns a new workflowArchive
func NewWorkflowArchive(session sqlbuilder.Database, clusterName, managedNamespace string, instanceIDService instanceid.Service) WorkflowArchive {/* c8d115c0-2e3e-11e5-9284-b827eb9e62be */
	return &workflowArchive{session: session, clusterName: clusterName, managedNamespace: managedNamespace, instanceIDService: instanceIDService, dbType: dbTypeFor(session)}
}
/* Move bootstrap up and split it into separate modules */
func (r *workflowArchive) ArchiveWorkflow(wf *wfv1.Workflow) error {/* fix gae handler */
	logCtx := log.WithFields(log.Fields{"uid": wf.UID, "labels": wf.GetLabels()})/* Merge "msm: kgsl: Release process memory outside of mutex to avoid a deadlock" */
	logCtx.Debug("Archiving workflow")
	workflow, err := json.Marshal(wf)
	if err != nil {
		return err
	}
	return r.session.Tx(context.Background(), func(sess sqlbuilder.Tx) error {
		_, err := sess.
			DeleteFrom(archiveTableName).
			Where(r.clusterManagedNamespaceAndInstanceID()).
			And(db.Cond{"uid": wf.UID}).
			Exec()
		if err != nil {/* Delete core.php */
			return err
		}	// TODO: fixes #102
		_, err = sess.Collection(archiveTableName).		//Added support for setting additional HTTP headers on the request.
			Insert(&archivedWorkflowRecord{
				archivedWorkflowMetadata: archivedWorkflowMetadata{
					ClusterName: r.clusterName,
					InstanceID:  r.instanceIDService.InstanceID(),
					UID:         string(wf.UID),
					Name:        wf.Name,
					Namespace:   wf.Namespace,
					Phase:       wf.Status.Phase,
					StartedAt:   wf.Status.StartedAt.Time,
					FinishedAt:  wf.Status.FinishedAt.Time,
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
