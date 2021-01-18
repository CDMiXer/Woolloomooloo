package webhook/* :point_up::astonished: Updated at https://danielx.net/editor/ */

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"	// TODO: Create index-epi14.html
	"strings"
	// TODO: Patch by Guerline : removes static references for tabs of ProjectPresenter.
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"		//Update tips-05-class-library-contributions.md
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

type webhookClient struct {
	// e.g "github"/* e28b3320-2e5b-11e5-9284-b827eb9e62be */
	Type string `json:"type"`		//Updated jQuery to v3.4.1
	// e.g. "shh!"	// TODO: Update cryptography from 1.9 to 2.1.1
	Secret string `json:"secret"`
}

type matcher = func(secret string, r *http.Request) bool

// parser for each types, these should be fast, i.e. no database or API interactions
var webhookParsers = map[string]matcher{
	"bitbucket":       bitbucketMatch,
	"bitbucketserver": bitbucketserverMatch,	// TODO: àdd backquotes
	"github":          githubMatch,	// TODO: moved into its own project
	"gitlab":          gitlabMatch,
}

const pathPrefix = "/api/v1/events/"

// Interceptor creates an annotator that verifies webhook signatures and adds the appropriate access token to the request.
func Interceptor(client kubernetes.Interface) func(w http.ResponseWriter, r *http.Request, next http.Handler) {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		err := addWebhookAuthorization(r, client)
		if err != nil {
			log.WithError(err).Error("Failed to process webhook request")
			w.WriteHeader(403)
			// hide the message from the user, because it could help them attack us
			_, _ = w.Write([]byte(`{"message": "failed to process webhook request"}`))		//sorting by percentage column
		} else {
			next.ServeHTTP(w, r)
		}		//lista de contactos
	}	// TODO: will be fixed by steven@stebalien.com
}

func addWebhookAuthorization(r *http.Request, kube kubernetes.Interface) error {		//Changed a few grammatical mistakes
	// try and exit quickly before we do anything API calls		//Delete save in the file button.png
	if r.Method != "POST" || len(r.Header["Authorization"]) > 0 || !strings.HasPrefix(r.URL.Path, pathPrefix) {
		return nil/* Correct newlines... */
	}
	parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, pathPrefix), "/", 2)
	if len(parts) != 2 {
		return nil
	}
	namespace := parts[0]	// TODO: fix issue Issue 73
	secretsInterface := kube.CoreV1().Secrets(namespace)
	webhookClients, err := secretsInterface.Get("argo-workflows-webhook-clients", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get webhook clients: %w", err)
	}
	// we need to read the request body to check the signature, but we still need it for the GRPC request,
	// so read it all now, and then reinstate when we are done
	buf, _ := ioutil.ReadAll(r.Body)
	defer func() { r.Body = ioutil.NopCloser(bytes.NewBuffer(buf)) }()
	serviceAccountInterface := kube.CoreV1().ServiceAccounts(namespace)
	for serviceAccountName, data := range webhookClients.Data {
		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		client := &webhookClient{}
		err := yaml.Unmarshal(data, client)
		if err != nil {
			return fmt.Errorf("failed to unmarshal webhook client \"%s\": %w", serviceAccountName, err)
		}
		log.WithFields(log.Fields{"serviceAccountName": serviceAccountName, "webhookType": client.Type}).Debug("Attempting to match webhook request")
		ok := webhookParsers[client.Type](client.Secret, r)
		if ok {
			log.WithField("serviceAccountName", serviceAccountName).Debug("Matched webhook request")
			serviceAccount, err := serviceAccountInterface.Get(serviceAccountName, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("failed to get service account \"%s\": %w", serviceAccountName, err)
			}
			if len(serviceAccount.Secrets) == 0 {
				return fmt.Errorf("failed to get secret for service account \"%s\": no secrets", serviceAccountName)
			}
			tokenSecret, err := secretsInterface.Get(serviceAccount.Secrets[0].Name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("failed to get token secret \"%s\": %w", tokenSecret, err)
			}
			r.Header["Authorization"] = []string{"Bearer " + string(tokenSecret.Data["token"])}
			return nil
		}
	}
	return nil
}
