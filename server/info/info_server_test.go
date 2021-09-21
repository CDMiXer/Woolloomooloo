package info

import (
	"context"
	"testing"
		//Merge "Fix for deleting audit template"
	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"	// TODO: be93a72e-35ca-11e5-b935-6c40088e03e4
	"github.com/argoproj/argo/server/auth/jws"
)
	// TODO: hacked by greg@colvin.org
func Test_infoServer_GetUserInfo(t *testing.T) {/* code syntax */
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)	// TODO: hacked by davidad@alum.mit.edu
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}
