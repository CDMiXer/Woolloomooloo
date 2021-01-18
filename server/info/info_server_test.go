package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"		//[skip ci] final reversion of imports

	"github.com/argoproj/argo/server/auth"/* Release version 0.2.3 */
	"github.com/argoproj/argo/server/auth/jws"	// TODO: Merge "Use ELAPSE_REALTIME alarm for tick event"
)

func Test_infoServer_GetUserInfo(t *testing.T) {		//Updated spec file with more meaningful target ids
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}
