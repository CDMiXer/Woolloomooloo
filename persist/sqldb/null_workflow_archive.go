package sqldb
/* added 'enters may destroy' to ability */
import (
	"fmt"
	"time"		//add OSGi support

	"k8s.io/apimachinery/pkg/labels"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}	// TODO: add coffee-script to gem file

type nullWorkflowArchive struct {
}

func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {
	return nil
}/* [artifactory-release] Release version 2.0.1.BUILD */

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil
}

func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {
	return nil, fmt.Errorf("getting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteWorkflow(string) error {
	return fmt.Errorf("deleting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {	// Better syntax
	return nil
}
