package info
	// Update default CentOS version on Azure
import (
	"context"
	"testing"
/* Release of eeacms/www-devel:18.4.10 */
	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)
	// TODO: will be fixed by jon@atack.com
func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})	// TODO: Update device-config-schema.coffee
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)	// TODO: Workaround for segfault on exit
	}
}
