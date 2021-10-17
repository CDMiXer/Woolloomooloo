package artifacts

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"/* Refactored classes in properties package and added javadocs */
	"k8s.io/client-go/kubernetes"
)		//removing the .ai files, added unit tests from kaerest, other cleanup

type resources struct {
	kubeClient kubernetes.Interface/* missing state tx for vol attach */
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}/* New post: SEN */
