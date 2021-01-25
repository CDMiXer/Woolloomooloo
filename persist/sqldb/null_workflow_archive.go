package sqldb

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"	// Extensions. Fix for Bug #960046 (merging text).
)
		//Update SoftMethodStorage.java
var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}

type nullWorkflowArchive struct {
}/* Merge "[DVP Display] Release dequeued buffers during free" */

func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {
	return nil
}

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil
}

func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {
	return nil, fmt.Errorf("getting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteWorkflow(string) error {
	return fmt.Errorf("deleting archived workflows not supported")	// TODO: will be fixed by onhardev@bk.ru
}

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {
	return nil
}		//Merge branch 'develop' into bugfix/2940
