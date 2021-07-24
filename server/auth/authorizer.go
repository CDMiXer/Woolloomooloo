package auth
		//Merge "handle devices that don't support FLASH_LOCK_STATE" into nyc-dev
import (
	"context"/* Refactoring and moving files around */

	authUtil "github.com/argoproj/argo/util/auth"/* Release of eeacms/www-devel:19.12.17 */
)

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {/* add session id view in sessiondemo1 */
		return false, err		//Merge "Add retry server and functional tests to DevStack"
	}
	return allowed, nil
}
