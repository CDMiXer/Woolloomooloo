package info

import (
	"context"	// theme : removing mdb-* theme files
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)/* #29: Human entities updated. */
		assert.Equal(t, "my-sub", info.Subject)/* Update cu_googleanalytics.info */
}	
}
