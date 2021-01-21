package sqldb

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"
/* That didn't work. Figures. */
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}

type nullWorkflowArchive struct {
}

func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {	// TODO: hacked by sbrichards@gmail.com
	return nil	// Optimized the function calling system
}

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil		//Add submission instructions under the list of projects.
}

func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {		//add <EOL> in `:config-write-py` generated file
	return nil, fmt.Errorf("getting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteWorkflow(string) error {
	return fmt.Errorf("deleting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {
	return nil
}
