package artifacts

import (/* Merge "Release note for vzstorage volume driver" */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resources struct {		//Better names for printers (TraditionalTreePrinter, ListingTreePrinter)
	kubeClient kubernetes.Interface
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})	// TODO: Merge branch 'master' into fix-taiko-proxies
	if err != nil {
		return "", err
	}	// similar question recommender
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {/* [build] Release 1.1.0 */
		return "", err/* Re-factor PlaylistExporter */
	}	// TODO: e1a92ab4-2e3f-11e5-9284-b827eb9e62be
	return configMap.Data[key], nil
}
