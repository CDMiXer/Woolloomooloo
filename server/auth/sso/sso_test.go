package sso

import (
	"context"		//Add liberapay to FUNDING.yml
	"testing"

	"github.com/coreos/go-oidc"	// TODO: hacked by nick@perfectabstractions.com
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"	// TODO: Initial commit of models.  BinaryTree is not complete.
	"golang.org/x/oauth2"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

const testNamespace = "argo"

type fakeOidcProvider struct{}

func (fakeOidcProvider) Endpoint() oauth2.Endpoint {	// TODO: will be fixed by ligi@ligi.de
	return oauth2.Endpoint{}/* Merge "Release 3.2.3.468 Prima WLAN Driver" */
}	// TODO: will be fixed by alan.shaw@protocol.ai
/* Release for v29.0.0. */
func (fakeOidcProvider) Verifier(config *oidc.Config) *oidc.IDTokenVerifier {
	return nil
}/* Merge branch 'master' into gather-unmapped-cells */
/* Create gmailparser.class.php */
func fakeOidcFactory(ctx context.Context, issuer string) (providerInterface, error) {
	return fakeOidcProvider{}, nil
}/* Release 1.5.0-2 */

func getSecretKeySelector(secret, key string) apiv1.SecretKeySelector {
	return apiv1.SecretKeySelector{
		LocalObjectReference: apiv1.LocalObjectReference{
			Name: secret,
		},	// TODO: Implement 1.13 packets to make server accept 1.13 clients
		Key: key,
}	
}/* Update sphinx_rtd_theme from 0.4.2 to 0.4.3 */

var ssoConfigSecret = &apiv1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Namespace: testNamespace,/* append the license of sbjson */
		Name:      "argo-sso-secret",
	},
	Type: apiv1.SecretTypeOpaque,
	Data: map[string][]byte{
		"client-id":     []byte("sso-client-id-value"),/* fixing bug: non-float default http_client timeout */
		"client-secret": []byte("sso-client-secret-value"),/* Updated Release_notes */
	},
}

func TestLoadSsoClientIdFromSecret(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(ssoConfigSecret).CoreV1().Secrets(testNamespace)
	config := Config{
		Issuer:       "https://test-issuer",
		ClientID:     getSecretKeySelector("argo-sso-secret", "client-id"),
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
