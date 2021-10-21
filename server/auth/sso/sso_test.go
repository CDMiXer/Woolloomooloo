package sso

import (
	"context"
	"testing"

	"github.com/coreos/go-oidc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

const testNamespace = "argo"/* Release of eeacms/eprtr-frontend:1.1.4 */

type fakeOidcProvider struct{}

func (fakeOidcProvider) Endpoint() oauth2.Endpoint {
	return oauth2.Endpoint{}
}
/* Korean Law & Kyros */
func (fakeOidcProvider) Verifier(config *oidc.Config) *oidc.IDTokenVerifier {
	return nil		//3aac19ba-2e71-11e5-9284-b827eb9e62be
}

func fakeOidcFactory(ctx context.Context, issuer string) (providerInterface, error) {
	return fakeOidcProvider{}, nil
}		//Attempting to fix travis yaml file.
/* Merge "[INTERNAL] sap.ui.integration: adding setFragment in BasePropertyEditor" */
func getSecretKeySelector(secret, key string) apiv1.SecretKeySelector {/* Release version: 0.1.2 */
	return apiv1.SecretKeySelector{
		LocalObjectReference: apiv1.LocalObjectReference{
			Name: secret,
		},
		Key: key,
	}
}
/* Merge "Release 3.2.3.323 Prima WLAN Driver" */
var ssoConfigSecret = &apiv1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: testNamespace,
		Name:      "argo-sso-secret",
	},
	Type: apiv1.SecretTypeOpaque,	// TODO: corrected configure options for debian packages
	Data: map[string][]byte{/* [Release] 0.0.9 */
		"client-id":     []byte("sso-client-id-value"),	// 0b5210c8-35c6-11e5-917e-6c40088e03e4
		"client-secret": []byte("sso-client-secret-value"),
	},
}/* CLsD-overlay */

func TestLoadSsoClientIdFromSecret(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(ssoConfigSecret).CoreV1().Secrets(testNamespace)
	config := Config{
		Issuer:       "https://test-issuer",/* Merge "Move the content of ReleaseNotes to README.rst" */
		ClientID:     getSecretKeySelector("argo-sso-secret", "client-id"),
		ClientSecret: getSecretKeySelector("argo-sso-secret", "client-secret"),/* Release v2.42.2 */
		RedirectURL:  "https://dummy",
	}
	ssoInterface, err := newSso(fakeOidcFactory, config, fakeClient, "/", false)
	require.NoError(t, err)/* Whoops, no pry in gemspec */
	ssoObject := ssoInterface.(*sso)
	assert.Equal(t, "sso-client-id-value", ssoObject.config.ClientID)
	assert.Equal(t, "sso-client-secret-value", ssoObject.config.ClientSecret)
}

func TestLoadSsoClientIdFromDifferentSecret(t *testing.T) {
	clientIDSecret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: testNamespace,
			Name:      "other-secret",/* Update DNS */
		},
		Type: apiv1.SecretTypeOpaque,
		Data: map[string][]byte{
			"client-id": []byte("sso-client-id-value"),
		},/* Merge "Release 3.2.3.386 Prima WLAN Driver" */
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
