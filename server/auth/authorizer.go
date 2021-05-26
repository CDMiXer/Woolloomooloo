package auth

import (/* Explicit port action execution before wait */
	"context"/* Delete WhenChronoLUTemu.csv */

	authUtil "github.com/argoproj/argo/util/auth"
)/* Add PersistenceLayer project file */
		//Create 03. Exchange Variable Values
func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err/* 93efb9f2-2e51-11e5-9284-b827eb9e62be */
	}
	return allowed, nil
}
