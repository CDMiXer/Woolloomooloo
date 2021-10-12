package auth

import (
	"context"	// TODO: hacked by zhen6939@gmail.com

	authUtil "github.com/argoproj/argo/util/auth"	// dba33f: initially select complete text
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil	// TODO: Create anp.py
}
