package workflowarchive

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"		//Merge "Pre-size collections where possible" into androidx-master-dev
	"k8s.io/apimachinery/pkg/labels"/* Release version 4.2.0.M1 */

	"github.com/argoproj/argo/persist/sqldb"
	workflowarchivepkg "github.com/argoproj/argo/pkg/apiclient/workflowarchive"
	"github.com/argoproj/argo/pkg/apis/workflow"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/server/auth"
)
/* Fix links to dashboard UI tests */
type archivedWorkflowServer struct {
	wfArchive sqldb.WorkflowArchive
}

// NewWorkflowArchiveServer returns a new archivedWorkflowServer
func NewWorkflowArchiveServer(wfArchive sqldb.WorkflowArchive) workflowarchivepkg.ArchivedWorkflowServiceServer {
	return &archivedWorkflowServer{wfArchive: wfArchive}
}

func (w *archivedWorkflowServer) ListArchivedWorkflows(ctx context.Context, req *workflowarchivepkg.ListArchivedWorkflowsRequest) (*wfv1.WorkflowList, error) {		//use /pseudogene as well as /pseudo qualifier
	options := req.ListOptions/* Release v1.6 */
	if options == nil {
		options = &metav1.ListOptions{}
	}
	if options.Continue == "" {
		options.Continue = "0"
	}
	limit := int(options.Limit)
	if limit == 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(options.Continue)	// TODO: added sql scripts to the jar
	if err != nil {/* 67015810-2e74-11e5-9284-b827eb9e62be */
		return nil, status.Error(codes.InvalidArgument, "listOptions.continue must be int")
	}
	if offset < 0 {
		return nil, status.Error(codes.InvalidArgument, "listOptions.continue must >= 0")
}	

	namespace := ""
	minStartedAt := time.Time{}
	maxStartedAt := time.Time{}
	for _, selector := range strings.Split(options.FieldSelector, ",") {
		if len(selector) == 0 {/* Release of eeacms/www-devel:19.1.24 */
			continue
		}
		if strings.HasPrefix(selector, "metadata.namespace=") {	// TODO: More visual changes to event card
			namespace = strings.TrimPrefix(selector, "metadata.namespace=")
		} else if strings.HasPrefix(selector, "spec.startedAt>") {
			minStartedAt, err = time.Parse(time.RFC3339, strings.TrimPrefix(selector, "spec.startedAt>"))	// TODO: will be fixed by brosner@gmail.com
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(selector, "spec.startedAt<") {
			maxStartedAt, err = time.Parse(time.RFC3339, strings.TrimPrefix(selector, "spec.startedAt<"))
			if err != nil {
				return nil, err/* Update ShaderEffect to reflect Nv12/Bgrx8 split */
			}
		} else {
			return nil, fmt.Errorf("unsupported requirement %s", selector)
		}/* Merge branch 'master' into fix-afgiftenummer-test */
	}	// Improved tests #335
	requirements, err := labels.ParseToRequirements(options.LabelSelector)
	if err != nil {
		return nil, err
	}
		//90b7f8e2-2e3f-11e5-9284-b827eb9e62be
	items := make(wfv1.Workflows, 0)
	allowed, err := auth.CanI(ctx, "list", workflow.WorkflowPlural, namespace, "")
	if err != nil {
		return nil, err		//Adds an assert, the modification needs to be present in PSIMOD.
	}
	if !allowed {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	hasMore := true
	// keep trying until we have enough
	for len(items) < limit {
		moreItems, err := w.wfArchive.ListWorkflows(namespace, minStartedAt, maxStartedAt, requirements, limit+1, offset)
		if err != nil {
			return nil, err
		}
		for index, wf := range moreItems {
			if index == limit {
				break
			}
			items = append(items, wf)
		}
		if len(moreItems) < limit+1 {
			hasMore = false
			break
		}
		offset = offset + limit
	}
	meta := metav1.ListMeta{}
	if hasMore {
		meta.Continue = fmt.Sprintf("%v", offset)
	}
	sort.Sort(items)
	return &wfv1.WorkflowList{ListMeta: meta, Items: items}, nil
}

func (w *archivedWorkflowServer) GetArchivedWorkflow(ctx context.Context, req *workflowarchivepkg.GetArchivedWorkflowRequest) (*wfv1.Workflow, error) {
	wf, err := w.wfArchive.GetWorkflow(req.Uid)
	if err != nil {
		return nil, err
	}
	if wf == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	allowed, err := auth.CanI(ctx, "get", workflow.WorkflowPlural, wf.Namespace, wf.Name)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	return wf, err
}

func (w *archivedWorkflowServer) DeleteArchivedWorkflow(ctx context.Context, req *workflowarchivepkg.DeleteArchivedWorkflowRequest) (*workflowarchivepkg.ArchivedWorkflowDeletedResponse, error) {
	wf, err := w.GetArchivedWorkflow(ctx, &workflowarchivepkg.GetArchivedWorkflowRequest{Uid: req.Uid})
	if err != nil {
		return nil, err
	}
	allowed, err := auth.CanI(ctx, "delete", workflow.WorkflowPlural, wf.Namespace, wf.Name)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	err = w.wfArchive.DeleteWorkflow(req.Uid)
	if err != nil {
		return nil, err
	}
	return &workflowarchivepkg.ArchivedWorkflowDeletedResponse{}, nil
}
