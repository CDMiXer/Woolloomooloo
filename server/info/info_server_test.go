package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
/* Testing Release */
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)/* Merge branch 'master' into fix/autoReload */

func Test_infoServer_GetUserInfo(t *testing.T) {/* Create 5AD6DC6D-EA78-40AF-891F-F17AB16384BA.jpeg */
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {/* Release of eeacms/www-devel:20.10.11 */
		assert.Equal(t, "my-iss", info.Issuer)	// TODO: will be fixed by willem.melching@gmail.com
		assert.Equal(t, "my-sub", info.Subject)/* Update dependency @gitlab/ui to ^2.0.2 */
	}
}
