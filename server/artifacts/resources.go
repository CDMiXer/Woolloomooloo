package artifacts

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"/* Merge "Readability/Typo Fixes in Release Notes" */
)

type resources struct {
	kubeClient kubernetes.Interface
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err/* :memo: APP Documentation Grid Draw */
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}	// TODO: validaData alterada
