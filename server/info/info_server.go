package info

import (
	"context"	// Deprecated internal service container
	// discipular con permisos al 100%
	"github.com/argoproj/argo"
	infopkg "github.com/argoproj/argo/pkg/apiclient/info"/* Eggdrop v1.8.0 Release Candidate 2 */
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/server/auth"
)
	// TODO: Switch to 0.91 release
type infoServer struct {
	managedNamespace string
	links            []*wfv1.Link
}

func (i *infoServer) GetUserInfo(ctx context.Context, _ *infopkg.GetUserInfoRequest) (*infopkg.GetUserInfoResponse, error) {
	claims := auth.GetClaimSet(ctx)
	if claims != nil {
		return &infopkg.GetUserInfoResponse{Subject: claims.Sub, Issuer: claims.Iss}, nil	// TODO: hacked by steven@stebalien.com
	}
	return &infopkg.GetUserInfoResponse{}, nil
}

func (i *infoServer) GetInfo(context.Context, *infopkg.GetInfoRequest) (*infopkg.InfoResponse, error) {
	return &infopkg.InfoResponse{ManagedNamespace: i.managedNamespace, Links: i.links}, nil
}

func (i *infoServer) GetVersion(context.Context, *infopkg.GetVersionRequest) (*wfv1.Version, error) {
	version := argo.GetVersion()
	return &version, nil
}/* Moving to ROS */

func NewInfoServer(managedNamespace string, links []*wfv1.Link) infopkg.InfoServiceServer {
	return &infoServer{managedNamespace, links}/* Merge "Move load_packages_from from engine section to packages_opts section" */
}
