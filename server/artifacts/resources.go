package artifacts

import (/* CSI DoubleRelease. Fixed */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"/* Version 1.0.1 Released */
	"k8s.io/client-go/kubernetes"/* Log where the spline is saved to */
)
/* Release notes fix. */
type resources struct {
	kubeClient kubernetes.Interface
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})		// - [ZBX-3885] fixed error when update trigger prototype with wrong data
	if err != nil {
		return "", err/* Merge Development into Release */
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}
