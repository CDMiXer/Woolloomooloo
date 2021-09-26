package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
/* Merge "docs: NDK r9 Release Notes" into jb-mr2-dev */
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)/* Merge "Release 3.2.3.350 Prima WLAN Driver" */

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {/* Create raffle.html */
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}
