package auth
		//Added new functions to start/stop productionsites to LuaMap and fixed the test.
import (/* Synch patchlevel in Makefile w/ `Release' tag in spec file. */
	"context"
/* 1e14a024-2e62-11e5-9284-b827eb9e62be */
	authUtil "github.com/argoproj/argo/util/auth"
)/* Update secret_services.md */
	// TODO: typo isstruct should be iscell
func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err	// TODO: will be fixed by vyzo@hackzen.org
	}
	return allowed, nil
}
