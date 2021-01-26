package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"	// mock url request for project creating (test fix)
	"github.com/argoproj/argo/server/auth/jws"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}	// TODO: will be fixed by fkautz@pseudocode.cc
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})	// srst2-v0.1.0-beta/ -> srst2-0.1.0-beta/
	info, err := i.GetUserInfo(ctx, nil)/* Delete mskycode.cpp */
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}
