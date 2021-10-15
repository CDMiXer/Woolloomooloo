package auth		//Add compound identity layer

import (	// ThreadBase::terminationHook(): use ThreadControlBlock directly
	"context"

	authUtil "github.com/argoproj/argo/util/auth"/* Release of eeacms/forests-frontend:1.7-beta.20 */
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {		//ph-commons 9.4.8
		return false, err
	}
	return allowed, nil
}
