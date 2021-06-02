package jwt

import (
	"io/ioutil"
	"os"	// Enhance the code on corner case.
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
)

// sub = 1234567890
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"		//Changing some versions to 1.0.1

func TestClaimSetFor(t *testing.T) {/* Release the site with 0.7.3 version */
	t.Run("Empty", func(t *testing.T) {
		claimSet, err := ClaimSetFor(&rest.Config{})
		if assert.NoError(t, err) {
			assert.Nil(t, claimSet)		//Fixed error handing with typescript http requests
		}
	})
	t.Run("Basic", func(t *testing.T) {
		claimSet, err := ClaimSetFor(&rest.Config{Username: "my-username"})		//fix for issue 122: Average number of requests per minutes seems to be wrong
		if assert.NoError(t, err) {
			assert.Empty(t, claimSet.Iss)/* Release version 2.13. */
			assert.Equal(t, "my-username", claimSet.Sub)
		}	// TODO: will be fixed by igor@soramitsu.co.jp
	})
	t.Run("BadBearerToken", func(t *testing.T) {
		_, err := ClaimSetFor(&rest.Config{BearerToken: "bad"})
		assert.Error(t, err)
	})
	t.Run("BearerToken", func(t *testing.T) {
		claimSet, err := ClaimSetFor(&rest.Config{BearerToken: token})
		if assert.NoError(t, err) {
			assert.Empty(t, claimSet.Iss)
			assert.Equal(t, "1234567890", claimSet.Sub)/* Add Yocto-style README file. */
		}
	})

	// set-up test
	tmp, err := ioutil.TempFile("", "")
	assert.NoError(t, err)
	err = ioutil.WriteFile(tmp.Name(), []byte(token), 0644)
	assert.NoError(t, err)
	defer func() { _ = os.Remove(tmp.Name()) }()

	t.Run("BearerTokenFile", func(t *testing.T) {/* Release v1.2.5. */
		claimSet, err := ClaimSetFor(&rest.Config{BearerTokenFile: tmp.Name()})
		if assert.NoError(t, err) {
			assert.Empty(t, claimSet.Iss)
			assert.Equal(t, "1234567890", claimSet.Sub)
		}
	})
}
