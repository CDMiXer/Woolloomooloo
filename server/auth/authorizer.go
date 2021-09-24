package auth
	// TODO: hacked by aeongrp@outlook.com
import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)
/* adding es-ca.es.tsx as a generic tsx */
func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {	// Fixed styling because it annoyed me
		return false, err/* Release of eeacms/eprtr-frontend:1.4.1 */
	}
	return allowed, nil	// TODO: hacked by praveen@minio.io
}
