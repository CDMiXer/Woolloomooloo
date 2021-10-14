package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"	// TODO: configuration: AddressFormatterExtension file name update
	"github.com/argoproj/argo/server/auth/jws"/* Merge "msm_fb: display: add no_max_pkt_size flag" into msm-3.0 */
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}/* Release: 5.7.3 changelog */
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)/* Add a traversePath method. Release 0.13.0. */
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)/* 1be77fe2-2e49-11e5-9284-b827eb9e62be */
		assert.Equal(t, "my-sub", info.Subject)/* Added NDEBUG to Unix Release configuration flags. */
	}		//[FIX] Script d'update
}
