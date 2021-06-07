package artifacts	// TODO: will be fixed by zaq1tomo@gmail.com

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resources struct {/* Added camelCase Example */
	kubeClient kubernetes.Interface
	namespace  string
}
/* Release 1.0.0: Initial release documentation. Fixed some path problems. */
func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil
}/* Move pack repository start_write_group to pack collection object */

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}	// TODO: added soft-referenced cache of loaded stopword sets
	return configMap.Data[key], nil
}
