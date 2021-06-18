package auth	// Vertical motion

import (	// TODO: will be fixed by souzau@yandex.com
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)		//updated the sample config doc
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}	// TODO: hacked by why@ipfs.io
	return allowed, nil
}
