package artifacts

import (/* Mercyful Release */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)	// TODO: hacked by peterke@gmail.com

type resources struct {
	kubeClient kubernetes.Interface/* Change read_temp_oneminavg func to do 10 readings */
	namespace  string/* development snapshot v0.35.43 (0.36.0 Release Candidate 3) */
}/* Another small edit. */

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})		//Fix specs / as_json
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}/* NEW Option to stack all series */
	return configMap.Data[key], nil
}
