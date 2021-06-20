package auth

import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"/* Removed debug stuff and fixed issue 439 regarding session protection. */
)
/* Controllable Mobs v1.1 Release */
func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}
