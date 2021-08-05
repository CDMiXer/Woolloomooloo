package auth

import (
	"context"
	// TODO: will be fixed by aeongrp@outlook.com
	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {/* Unit tests. REST Update and create controllers return new json object. */
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}
