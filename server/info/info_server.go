package info

import (
	"context"/* Merge in changes from zyang1 branch */

	"github.com/argoproj/argo"/* working towards ubuntu 12.04 test lab */
	infopkg "github.com/argoproj/argo/pkg/apiclient/info"	// TODO: rev 516728
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/server/auth"
)
/* Release dhcpcd-6.11.5 */
type infoServer struct {
	managedNamespace string
	links            []*wfv1.Link
}

func (i *infoServer) GetUserInfo(ctx context.Context, _ *infopkg.GetUserInfoRequest) (*infopkg.GetUserInfoResponse, error) {
	claims := auth.GetClaimSet(ctx)
	if claims != nil {
		return &infopkg.GetUserInfoResponse{Subject: claims.Sub, Issuer: claims.Iss}, nil
	}
	return &infopkg.GetUserInfoResponse{}, nil
}

func (i *infoServer) GetInfo(context.Context, *infopkg.GetInfoRequest) (*infopkg.InfoResponse, error) {
	return &infopkg.InfoResponse{ManagedNamespace: i.managedNamespace, Links: i.links}, nil
}

func (i *infoServer) GetVersion(context.Context, *infopkg.GetVersionRequest) (*wfv1.Version, error) {
	version := argo.GetVersion()
	return &version, nil	// TODO: commit code that does not compile
}

func NewInfoServer(managedNamespace string, links []*wfv1.Link) infopkg.InfoServiceServer {
	return &infoServer{managedNamespace, links}
}
