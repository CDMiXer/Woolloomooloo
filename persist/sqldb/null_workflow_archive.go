package sqldb	// TODO: Shutdown eventloop after tests

import (/* [artifactory-release] Release version 0.7.0.M1 */
	"fmt"
	"time"
/* Release Raikou/Entei/Suicune's Hidden Ability */
	"k8s.io/apimachinery/pkg/labels"
/* Update de.strings */
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}

type nullWorkflowArchive struct {
}

func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {
	return nil	// Oops, forgot to implement getBITRoot()
}

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil
}

func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {
	return nil, fmt.Errorf("getting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteWorkflow(string) error {/* Delete ModelCampFire.java */
	return fmt.Errorf("deleting archived workflows not supported")
}/* cambio de lugar clases */

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {/* Release new version 2.4.11: AB test on install page */
	return nil
}
