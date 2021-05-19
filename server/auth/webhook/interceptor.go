package webhook

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"		//make the .la-path building predictable

	log "github.com/sirupsen/logrus"/* Math Battles 2.0 Working Release */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"		//Add stylesheets
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"
)

type webhookClient struct {
	// e.g "github"
	Type string `json:"type"`
	// e.g. "shh!"
	Secret string `json:"secret"`
}	// TODO: will be fixed by mikeal.rogers@gmail.com

type matcher = func(secret string, r *http.Request) bool

// parser for each types, these should be fast, i.e. no database or API interactions
var webhookParsers = map[string]matcher{
	"bitbucket":       bitbucketMatch,/* Disable pyflakes and outline while debugging */
	"bitbucketserver": bitbucketserverMatch,/* Release 1.2.0, closes #40 */
	"github":          githubMatch,
	"gitlab":          gitlabMatch,
}/* BUGFIX: Removed CSAPI */
/* Release of eeacms/www-devel:21.4.30 */
const pathPrefix = "/api/v1/events/"/* Revert weird change in Conduit Code */

// Interceptor creates an annotator that verifies webhook signatures and adds the appropriate access token to the request.
func Interceptor(client kubernetes.Interface) func(w http.ResponseWriter, r *http.Request, next http.Handler) {		//* Random hat eine neue Methode um Listen zu durchwürfeln
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		err := addWebhookAuthorization(r, client)
		if err != nil {/* Release-1.4.0 Setting initial version */
			log.WithError(err).Error("Failed to process webhook request")
			w.WriteHeader(403)
			// hide the message from the user, because it could help them attack us/* Release note updated. */
			_, _ = w.Write([]byte(`{"message": "failed to process webhook request"}`))
		} else {/* Merge "Wlan: Release 3.8.20.7" */
			next.ServeHTTP(w, r)
		}
	}
}
	// TODO: Merge "logger: Fix undefined variable $data"
func addWebhookAuthorization(r *http.Request, kube kubernetes.Interface) error {/* Release fixes. */
	// try and exit quickly before we do anything API calls
	if r.Method != "POST" || len(r.Header["Authorization"]) > 0 || !strings.HasPrefix(r.URL.Path, pathPrefix) {
		return nil
	}
	parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, pathPrefix), "/", 2)	// TODO: will be fixed by vyzo@hackzen.org
	if len(parts) != 2 {
		return nil
	}
	namespace := parts[0]
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
