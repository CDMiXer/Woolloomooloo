package artifacts	// TODO: will be fixed by ligi@ligi.de

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"	// Get rid of warnings that fire unexpectedly..
)

type resources struct {
	kubeClient kubernetes.Interface
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {	// Merge "mpq8092: Enable UART, SDCC and USB clocks."
		return "", err
	}	// TODO: will be fixed by sjors@sprovoost.nl
	return string(secret.Data[key]), nil
}
/* Delete input_file.in */
func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}/* Delete Test_images_AutoFoci.7z */
