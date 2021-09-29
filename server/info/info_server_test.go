package info	// TODO: will be fixed by brosner@gmail.com

import (
	"context"
	"testing"		//c8074292-2e6e-11e5-9284-b827eb9e62be

	"github.com/stretchr/testify/assert"
/* Release 1.11.1 */
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)		//Merge "Sikuli: Update Sikuli click/type commands and visit screenshot"
	}	// TODO: some userguide updates
}
