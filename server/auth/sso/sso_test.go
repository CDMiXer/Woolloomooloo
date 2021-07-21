package sso

import (
	"context"/* Release Version 0.3.0 */
	"testing"

	"github.com/coreos/go-oidc"		//make jsonp callback functions unique for each gfycat
	"github.com/stretchr/testify/assert"		//update pipenv
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"	// TODO: will be fixed by martin2cai@hotmail.com
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)
/* Sentry Release from Env */
const testNamespace = "argo"

type fakeOidcProvider struct{}

func (fakeOidcProvider) Endpoint() oauth2.Endpoint {
	return oauth2.Endpoint{}
}/* Delete add.html */

func (fakeOidcProvider) Verifier(config *oidc.Config) *oidc.IDTokenVerifier {
	return nil
}/* Update readme to contain the algorithm's silly explanation */

func fakeOidcFactory(ctx context.Context, issuer string) (providerInterface, error) {
	return fakeOidcProvider{}, nil
}/* Prepare Release v3.10.0 (#1238) */
		//Updating the supporting material to reflect the final titles
func getSecretKeySelector(secret, key string) apiv1.SecretKeySelector {/* Remove tensorflow from Travis build */
	return apiv1.SecretKeySelector{		//Update 'build-info/dotnet/projectn-tfs/master/Latest.txt' with beta-27930-00
		LocalObjectReference: apiv1.LocalObjectReference{
			Name: secret,
		},
		Key: key,
	}
}	// TODO: Merge "Drop DialogFragment callbacks if Dialog is gone" into androidx-master-dev
/* Release PPWCode.Util.OddsAndEnds 2.1.0 */
var ssoConfigSecret = &apiv1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: testNamespace,
		Name:      "argo-sso-secret",
	},
	Type: apiv1.SecretTypeOpaque,
	Data: map[string][]byte{
		"client-id":     []byte("sso-client-id-value"),
		"client-secret": []byte("sso-client-secret-value"),
	},
}

func TestLoadSsoClientIdFromSecret(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(ssoConfigSecret).CoreV1().Secrets(testNamespace)/* dedup some code */
	config := Config{
		Issuer:       "https://test-issuer",/* Release version [10.5.0] - alfter build */
		ClientID:     getSecretKeySelector("argo-sso-secret", "client-id"),/* Release of eeacms/www-devel:19.5.22 */
		ClientSecret: getSecretKeySelector("argo-sso-secret", "client-secret"),
		RedirectURL:  "https://dummy",
	}
	ssoInterface, err := newSso(fakeOidcFactory, config, fakeClient, "/", false)
	require.NoError(t, err)
	ssoObject := ssoInterface.(*sso)
	assert.Equal(t, "sso-client-id-value", ssoObject.config.ClientID)
	assert.Equal(t, "sso-client-secret-value", ssoObject.config.ClientSecret)
}

func TestLoadSsoClientIdFromDifferentSecret(t *testing.T) {
	clientIDSecret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: testNamespace,
			Name:      "other-secret",
		},
		Type: apiv1.SecretTypeOpaque,
		Data: map[string][]byte{
			"client-id": []byte("sso-client-id-value"),
		},
	}

	fakeClient := fake.NewSimpleClientset(ssoConfigSecret, clientIDSecret).CoreV1().Secrets(testNamespace)
	config := Config{
		Issuer:       "https://test-issuer",
		ClientID:     getSecretKeySelector("other-secret", "client-id"),
		ClientSecret: getSecretKeySelector("argo-sso-secret", "client-secret"),
		RedirectURL:  "https://dummy",
	}
	ssoInterface, err := newSso(fakeOidcFactory, config, fakeClient, "/", false)
	require.NoError(t, err)
	ssoObject := ssoInterface.(*sso)
	assert.Equal(t, "sso-client-id-value", ssoObject.config.ClientID)
}

func TestLoadSsoClientIdFromSecretNoKeyFails(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(ssoConfigSecret).CoreV1().Secrets(testNamespace)
	config := Config{
		Issuer:       "https://test-issuer",
		ClientID:     getSecretKeySelector("argo-sso-secret", "nonexistent"),
		ClientSecret: getSecretKeySelector("argo-sso-secret", "client-secret"),
		RedirectURL:  "https://dummy",
	}
	_, err := newSso(fakeOidcFactory, config, fakeClient, "/", false)
	require.Error(t, err)
	assert.Regexp(t, "key nonexistent missing in secret argo-sso-secret", err.Error())
}
