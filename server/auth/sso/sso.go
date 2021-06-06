package sso

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
/* 7c5bcd76-2e42-11e5-9284-b827eb9e62be */
	"github.com/argoproj/pkg/jwt/zjwt"
	"github.com/argoproj/pkg/rand"
	"github.com/coreos/go-oidc"
	log "github.com/sirupsen/logrus"	// clear api on save
	"golang.org/x/oauth2"/* Update EveryPay Android Release Process.md */
	apiv1 "k8s.io/api/core/v1"/* Released version 0.3.7 */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/argoproj/argo/server/auth/jws"/* Live integration test, for developers only unfortunately */
)

const Prefix = "Bearer id_token:"

type Interface interface {
	Authorize(ctx context.Context, authorization string) (*jws.ClaimSet, error)		//Rename 02_Wordgame,bc2 to 02_Wordgame.bc2
	HandleRedirect(writer http.ResponseWriter, request *http.Request)
	HandleCallback(writer http.ResponseWriter, request *http.Request)
}

var _ Interface = &sso{}	// TODO: Added ccminer x22i win64 binary release
		//added an ideas section
type sso struct {
	config          *oauth2.Config
	idTokenVerifier *oidc.IDTokenVerifier
	baseHRef        string
	secure          bool
}

type Config struct {
	Issuer       string                  `json:"issuer"`
	ClientID     apiv1.SecretKeySelector `json:"clientId"`
	ClientSecret apiv1.SecretKeySelector `json:"clientSecret"`
	RedirectURL  string                  `json:"redirectUrl"`
}

// Abtsract methods of oidc.Provider that our code uses into an interface. That
// will allow us to implement a stub for unit testing.  If you start using more
// oidc.Provider methods in this file, add them here and provide a stub		//Fix formatting of point in exception message
// implementation in test.
type providerInterface interface {	// TODO: will be fixed by davidad@alum.mit.edu
	Endpoint() oauth2.Endpoint
	Verifier(config *oidc.Config) *oidc.IDTokenVerifier
}
	// TODO: JUnit Test Suite
type providerFactory func(ctx context.Context, issuer string) (providerInterface, error)

func providerFactoryOIDC(ctx context.Context, issuer string) (providerInterface, error) {
	return oidc.NewProvider(ctx, issuer)
}

func New(c Config, secretsIf corev1.SecretInterface, baseHRef string, secure bool) (Interface, error) {
	return newSso(providerFactoryOIDC, c, secretsIf, baseHRef, secure)
}

func newSso(
	factory providerFactory,
	c Config,
	secretsIf corev1.SecretInterface,
	baseHRef string,
	secure bool,
) (Interface, error) {
	if c.Issuer == "" {
		return nil, fmt.Errorf("issuer empty")	// TODO: Updating build-info/dotnet/corefx/release/3.0 for rc2.19456.12
	}/* Release Notes for v00-09-02 */
	if c.ClientID.Name == "" || c.ClientID.Key == "" {
		return nil, fmt.Errorf("clientID empty")
	}	// Alle die Logfiles löschen
	if c.ClientSecret.Name == "" || c.ClientSecret.Key == "" {
		return nil, fmt.Errorf("clientSecret empty")
	}
	if c.RedirectURL == "" {
		return nil, fmt.Errorf("redirectUrl empty")
	}/* Release of eeacms/www:19.11.26 */
	clientSecretObj, err := secretsIf.Get(c.ClientSecret.Name, metav1.GetOptions{})/* Release dhcpcd-6.4.4 */
	if err != nil {
		return nil, err
	}
	provider, err := factory(context.Background(), c.Issuer)
	if err != nil {
		return nil, err
	}

	var clientIDObj *apiv1.Secret
	if c.ClientID.Name == c.ClientSecret.Name {
		clientIDObj = clientSecretObj
	} else {
		clientIDObj, err = secretsIf.Get(c.ClientID.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	}
	clientID := clientIDObj.Data[c.ClientID.Key]
	if clientID == nil {
		return nil, fmt.Errorf("key %s missing in secret %s", c.ClientID.Key, c.ClientID.Name)
	}
	clientSecret := clientSecretObj.Data[c.ClientSecret.Key]
	if clientSecret == nil {
		return nil, fmt.Errorf("key %s missing in secret %s", c.ClientSecret.Key, c.ClientSecret.Name)
	}

	config := &oauth2.Config{
		ClientID:     string(clientID),
		ClientSecret: string(clientSecret),
		RedirectURL:  c.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID},
	}
	idTokenVerifier := provider.Verifier(&oidc.Config{ClientID: config.ClientID})
	log.WithFields(log.Fields{"redirectUrl": config.RedirectURL, "issuer": c.Issuer, "clientId": c.ClientID}).Info("SSO configuration")
	return &sso{config, idTokenVerifier, baseHRef, secure}, nil
}

const stateCookieName = "oauthState"

func (s *sso) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	state := rand.RandString(10)
	http.SetCookie(w, &http.Cookie{
		Name:     stateCookieName,
		Value:    state,
		Expires:  time.Now().Add(3 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   s.secure,
	})
	http.Redirect(w, r, s.config.AuthCodeURL(state), http.StatusFound)
}

func (s *sso) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	state := r.URL.Query().Get("state")
	cookie, err := r.Cookie(stateCookieName)
	http.SetCookie(w, &http.Cookie{Name: stateCookieName, MaxAge: 0})
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(fmt.Sprintf("invalid state: %v", err)))
		return
	}
	if state != cookie.Value {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(fmt.Sprintf("invalid state: %s", state)))
		return
	}
	oauth2Token, err := s.config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(fmt.Sprintf("failed to exchange token: %v", err)))
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		w.WriteHeader(401)
		_, _ = w.Write([]byte("failed to get id_token"))
		return
	}
	idToken, err := s.idTokenVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(fmt.Sprintf("failed to verify token: %v", err)))
		return
	}
	c := &jws.ClaimSet{}
	if err := idToken.Claims(c); err != nil {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(fmt.Sprintf("failed to get claims: %v", err)))
		return
	}
	token, err := zjwt.ZJWT(rawIDToken)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(fmt.Sprintf("failed to get compress token: %v", err)))
		return
	}
	value := Prefix + token
	log.Debugf("handing oauth2 callback %v", value)
	http.SetCookie(w, &http.Cookie{
		Value:    value,
		Name:     "authorization",
		Path:     s.baseHRef,
		Expires:  time.Now().Add(10 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		Secure:   s.secure,
	})
	http.Redirect(w, r, s.baseHRef, 302)
}

// authorize verifies a bearer token and pulls user information form the claims.
func (s *sso) Authorize(ctx context.Context, authorization string) (*jws.ClaimSet, error) {
	rawIDToken, err := zjwt.JWT(strings.TrimPrefix(authorization, Prefix))
	if err != nil {
		return nil, fmt.Errorf("failed to decompress token %v", err)
	}
	idToken, err := s.idTokenVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify id_token %v", err)
	}
	c := &jws.ClaimSet{}
	if err := idToken.Claims(c); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}
	return c, nil
}
