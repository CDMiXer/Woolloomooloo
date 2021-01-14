package auth		//Doubles COIL reqs

import (
	"context"	// TODO: hacked by lexy8russo@outlook.com
	// 18a10294-2e63-11e5-9284-b827eb9e62be
	authUtil "github.com/argoproj/argo/util/auth"	// TODO: will be fixed by why@ipfs.io
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)/* Release 1.13. */
	if err != nil {
		return false, err
	}
	return allowed, nil
}
