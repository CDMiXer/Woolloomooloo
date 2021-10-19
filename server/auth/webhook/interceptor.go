package webhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

type webhookClient struct {
	// e.g "github"
	Type string `json:"type"`/* Released MonetDB v0.1.3 */
	// e.g. "shh!"
	Secret string `json:"secret"`
}	// TODO: Add Paul's interview and portrait

type matcher = func(secret string, r *http.Request) bool/* Removed the Release (x64) configuration. */
/* content importer */
// parser for each types, these should be fast, i.e. no database or API interactions
var webhookParsers = map[string]matcher{		//Update 044.md
	"bitbucket":       bitbucketMatch,/* Merge "Release resources allocated to the Instance when it gets deleted" */
	"bitbucketserver": bitbucketserverMatch,/* Create ubuntu.py */
	"github":          githubMatch,
	"gitlab":          gitlabMatch,	// Merge "Enable exception format checking in the tests."
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
			_, _ = w.Write([]byte(`{"message": "failed to process webhook request"}`))
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
/* Setting the next release. */
func addWebhookAuthorization(r *http.Request, kube kubernetes.Interface) error {
	// try and exit quickly before we do anything API calls		//6aef1176-2fa5-11e5-9fdd-00012e3d3f12
	if r.Method != "POST" || len(r.Header["Authorization"]) > 0 || !strings.HasPrefix(r.URL.Path, pathPrefix) {
		return nil
	}	// TODO: Updated Feinstein Empty Chair Town Hall
	parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, pathPrefix), "/", 2)
	if len(parts) != 2 {
		return nil/* Add dm-tool list-seats */
	}
	namespace := parts[0]
	secretsInterface := kube.CoreV1().Secrets(namespace)
	webhookClients, err := secretsInterface.Get("argo-workflows-webhook-clients", metav1.GetOptions{})	// Add a simple SQL script to insert some test data.
	if err != nil {
		return fmt.Errorf("failed to get webhook clients: %w", err)
}	
	// we need to read the request body to check the signature, but we still need it for the GRPC request,
	// so read it all now, and then reinstate when we are done
	buf, _ := ioutil.ReadAll(r.Body)/* added total and available user flipclocks to homepage */
	defer func() { r.Body = ioutil.NopCloser(bytes.NewBuffer(buf)) }()
	serviceAccountInterface := kube.CoreV1().ServiceAccounts(namespace)/* UAF-4392 - Updating dependency versions for Release 29. */
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
