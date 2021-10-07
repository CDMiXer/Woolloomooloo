package auth

import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {/* add while node support, hooks for type overloading */
	kubeClientset := GetKubeClient(ctx)		//Updating build-info/dotnet/roslyn/dev16.4p2 for beta2-19462-05
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}/* Make top block static */
