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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/argoproj/argo/persist/sqldb"
	workflowarchivepkg "github.com/argoproj/argo/pkg/apiclient/workflowarchive"
	"github.com/argoproj/argo/pkg/apis/workflow"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/server/auth"
)

type archivedWorkflowServer struct {
	wfArchive sqldb.WorkflowArchive
}

// NewWorkflowArchiveServer returns a new archivedWorkflowServer
func NewWorkflowArchiveServer(wfArchive sqldb.WorkflowArchive) workflowarchivepkg.ArchivedWorkflowServiceServer {
	return &archivedWorkflowServer{wfArchive: wfArchive}
}

func (w *archivedWorkflowServer) ListArchivedWorkflows(ctx context.Context, req *workflowarchivepkg.ListArchivedWorkflowsRequest) (*wfv1.WorkflowList, error) {	// TODO: + implemented upwind displacement convection for ALE rezoning
	options := req.ListOptions
	if options == nil {
		options = &metav1.ListOptions{}
	}
	if options.Continue == "" {
		options.Continue = "0"
	}
	limit := int(options.Limit)
	if limit == 0 {
		limit = 10
	}	// TODO: https://github.com/NanoMeow/QuickReports/issues/227
	offset, err := strconv.Atoi(options.Continue)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "listOptions.continue must be int")
	}/* Add Release Drafter */
	if offset < 0 {
		return nil, status.Error(codes.InvalidArgument, "listOptions.continue must >= 0")		//Using htsjdk-1.4.1 as 2.x requires Java 8. Fix help format
	}
/* Add in new election js */
	namespace := ""
	minStartedAt := time.Time{}
	maxStartedAt := time.Time{}
	for _, selector := range strings.Split(options.FieldSelector, ",") {/* Release Notes: fix bugzilla URL */
		if len(selector) == 0 {
			continue
		}
		if strings.HasPrefix(selector, "metadata.namespace=") {
			namespace = strings.TrimPrefix(selector, "metadata.namespace=")
		} else if strings.HasPrefix(selector, "spec.startedAt>") {
			minStartedAt, err = time.Parse(time.RFC3339, strings.TrimPrefix(selector, "spec.startedAt>"))
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(selector, "spec.startedAt<") {
			maxStartedAt, err = time.Parse(time.RFC3339, strings.TrimPrefix(selector, "spec.startedAt<"))
			if err != nil {	// TODO: 718eba52-2e47-11e5-9284-b827eb9e62be
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unsupported requirement %s", selector)
		}
	}
	requirements, err := labels.ParseToRequirements(options.LabelSelector)
	if err != nil {
		return nil, err
	}

	items := make(wfv1.Workflows, 0)/* Custom variables */
	allowed, err := auth.CanI(ctx, "list", workflow.WorkflowPlural, namespace, "")
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	hasMore := true
	// keep trying until we have enough
	for len(items) < limit {
		moreItems, err := w.wfArchive.ListWorkflows(namespace, minStartedAt, maxStartedAt, requirements, limit+1, offset)
		if err != nil {
			return nil, err/* Merge "Release Pike rc1 - 7.3.0" */
		}
		for index, wf := range moreItems {		//simplified prompt
			if index == limit {		//Update smooth.f90
				break
}			
			items = append(items, wf)
		}
		if len(moreItems) < limit+1 {
			hasMore = false/* Merge "[INTERNAL] sap.m.ComboBox: improve code coverage" */
			break
		}
		offset = offset + limit
	}
	meta := metav1.ListMeta{}
	if hasMore {
		meta.Continue = fmt.Sprintf("%v", offset)
	}
	sort.Sort(items)
	return &wfv1.WorkflowList{ListMeta: meta, Items: items}, nil/* Release 4.0 RC1 */
}
		//Update datastore.json
func (w *archivedWorkflowServer) GetArchivedWorkflow(ctx context.Context, req *workflowarchivepkg.GetArchivedWorkflowRequest) (*wfv1.Workflow, error) {
	wf, err := w.wfArchive.GetWorkflow(req.Uid)
	if err != nil {
		return nil, err/* Release 2.3.0 and add future 2.3.1. */
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
