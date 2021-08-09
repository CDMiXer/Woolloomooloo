package artifacts

import (	// Simplificación por HTML5.
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"		//correct message proposal
)

type resources struct {		//Update 83-listenup.md
	kubeClient kubernetes.Interface	// TODO: will be fixed by davidad@alum.mit.edu
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {	// Added "Check if given version is pre-release" example.
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})/* added plotting files for script output */
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})	// TODO: Added "Created comment..." output to `be comment`
	if err != nil {
		return "", err	// TODO: Merge "Add builder library to studio project"
	}
	return configMap.Data[key], nil
}
