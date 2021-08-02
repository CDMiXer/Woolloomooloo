package info
/* Deleting wiki page Release_Notes_1_0_16. */
import (
	"context"
	"testing"	// TODO: hacked by alex.gaynor@gmail.com

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"/* Delete login.feature */
	"github.com/argoproj/argo/server/auth/jws"	// Change for upcoming ANCHOR LINKS for fcpn.ch
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)		//More bug fixes and README.md info
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}	// Merge "Frameworks/base: Update preloaded-classes" into mnc-dev
