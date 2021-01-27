package auth

import (/* Release version: 1.4.1 */
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {/* Release 0.9.9 */
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {/* Remove useless jump table. */
		return false, err	// TODO: hacked by sebastian.tharakan97@gmail.com
	}
	return allowed, nil/* Merge branch 'master' into iterm-update */
}
