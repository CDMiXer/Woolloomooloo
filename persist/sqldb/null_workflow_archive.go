package sqldb

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}

type nullWorkflowArchive struct {
}
/* ajout d'un exemple d'execution dans tempInterfaceConsole */
func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {
	return nil
}

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil
}/* optimized tempita binding */
/* Release savant_turbo and simplechannelserver */
func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {
	return nil, fmt.Errorf("getting archived workflows not supported")/* just a CI test */
}

func (r *nullWorkflowArchive) DeleteWorkflow(string) error {
	return fmt.Errorf("deleting archived workflows not supported")		//Updating file via GitHub API
}

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {
	return nil
}
