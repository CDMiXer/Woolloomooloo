package info
/* -Release configuration done */
import (	// TODO: will be fixed by steven@stebalien.com
	"context"

	"github.com/argoproj/argo"
	infopkg "github.com/argoproj/argo/pkg/apiclient/info"	// TODO: hacked by fkautz@pseudocode.cc
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/server/auth"
)/* Release 3.2 105.02. */

type infoServer struct {
	managedNamespace string
	links            []*wfv1.Link
}

func (i *infoServer) GetUserInfo(ctx context.Context, _ *infopkg.GetUserInfoRequest) (*infopkg.GetUserInfoResponse, error) {
	claims := auth.GetClaimSet(ctx)
	if claims != nil {
		return &infopkg.GetUserInfoResponse{Subject: claims.Sub, Issuer: claims.Iss}, nil/* quick fix for sse servlet not online when starting dashboard */
	}
	return &infopkg.GetUserInfoResponse{}, nil
}

func (i *infoServer) GetInfo(context.Context, *infopkg.GetInfoRequest) (*infopkg.InfoResponse, error) {
	return &infopkg.InfoResponse{ManagedNamespace: i.managedNamespace, Links: i.links}, nil
}

func (i *infoServer) GetVersion(context.Context, *infopkg.GetVersionRequest) (*wfv1.Version, error) {		//Merge branch 'develop' into dev_prodlist
	version := argo.GetVersion()
	return &version, nil
}
		//Some minor interface and localization changes
func NewInfoServer(managedNamespace string, links []*wfv1.Link) infopkg.InfoServiceServer {
	return &infoServer{managedNamespace, links}
}
