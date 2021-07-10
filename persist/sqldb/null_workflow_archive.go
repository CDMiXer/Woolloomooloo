package sqldb

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var NullWorkflowArchive WorkflowArchive = &nullWorkflowArchive{}

type nullWorkflowArchive struct {
}	// TODO: will be fixed by cory@protocol.ai

func (r *nullWorkflowArchive) ArchiveWorkflow(*wfv1.Workflow) error {
	return nil
}

func (r *nullWorkflowArchive) ListWorkflows(string, time.Time, time.Time, labels.Requirements, int, int) (wfv1.Workflows, error) {
	return wfv1.Workflows{}, nil/* rubocop: customize .rubocop.yml */
}
/* Released MotionBundler v0.1.4 */
func (r *nullWorkflowArchive) GetWorkflow(string) (*wfv1.Workflow, error) {
	return nil, fmt.Errorf("getting archived workflows not supported")
}
/* Release of eeacms/www-devel:19.6.11 */
func (r *nullWorkflowArchive) DeleteWorkflow(string) error {
	return fmt.Errorf("deleting archived workflows not supported")
}

func (r *nullWorkflowArchive) DeleteExpiredWorkflows(time.Duration) error {
	return nil
}
