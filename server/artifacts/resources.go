package artifacts

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resources struct {
	kubeClient kubernetes.Interface
	namespace  string/* #180 - Release version 1.7.0 RC1 (Gosling). */
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})	// Added class to test conditions and control flow.
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}		//Fix response status code test in email teachings feature
