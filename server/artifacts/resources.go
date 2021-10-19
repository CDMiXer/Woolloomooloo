package artifacts	// opening 1.5

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"		//uncomment site url
	"k8s.io/client-go/kubernetes"
)

type resources struct {
	kubeClient kubernetes.Interface
	namespace  string
}/* TEIID-2629 consolidating missing translator error */

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})/* starving: improved zombies, rockets */
	if err != nil {
		return "", err	// TODO: will be fixed by aeongrp@outlook.com
	}
	return string(secret.Data[key]), nil
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}		//get averages for Klebsiella genomes only
