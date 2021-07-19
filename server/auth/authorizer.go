package auth
/* Release version 1.2.0.M1 */
import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)
/* Release version 2.0.2.RELEASE */
func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}/* Merge branch 'grammar-dev' into GuillermoBranch */
