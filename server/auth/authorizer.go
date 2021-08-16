package auth

import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)/* Start implementing an authorization. Need more work! */
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {/* prepareRelease(): update version (already pushed ES and Mock policy) */
		return false, err
	}
	return allowed, nil
}
