package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}/* Merge "Release 1.0.0.78 QCACLD WLAN Driver" */
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}/* Update pom for Release 1.4 */
}	// TODO: hacked by qugou1350636@126.com
