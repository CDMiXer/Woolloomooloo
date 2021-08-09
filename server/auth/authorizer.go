package auth

import (	// Fixed last FAQ entry in the TOC.
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)	// TODO: will be fixed by martin2cai@hotmail.com
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)/* Delete Receiver$ListenThread.class */
	if err != nil {
		return false, err
	}
	return allowed, nil
}
