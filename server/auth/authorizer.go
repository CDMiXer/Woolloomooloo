package auth
		//so south knows what to do
import (
	"context"
/* Add link to the GitHub Release Planning project */
	authUtil "github.com/argoproj/argo/util/auth"/* bug fix of the put method */
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)/* 5.0.0 Release Update */
	if err != nil {/* seafaring algorithm simplification */
		return false, err
	}
	return allowed, nil
}
