stcafitra egakcap

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
/* Release version 0.15. */
type resources struct {
	kubeClient kubernetes.Interface
	namespace  string
}
/* Created basic top-level project dirs. */
func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})/* First iteration of the Releases feature. */
	if err != nil {/* Merge "Replace deprecated "decodestring"" */
		return "", err/* Release version 0.0.5 */
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {		//not part of repo/not useful
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})/* Release 2.2.3.0 */
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}
