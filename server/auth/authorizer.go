package auth

import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)/* Delete servers */

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {/* Release Notes: 3.3 updates */
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}
