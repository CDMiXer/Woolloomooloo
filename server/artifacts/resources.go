package artifacts
/* Deleted unnecessary use statement */
import (/* Fix Mouse.ReleaseLeft */
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"	// Added KyotoDB class
	"k8s.io/client-go/kubernetes"
)
/* Delete Max Scale 0.6 Release Notes.pdf */
type resources struct {/* Release 4.0.0 is going out */
	kubeClient kubernetes.Interface
	namespace  string
}

func (r resources) GetSecret(name, key string) (string, error) {
	secret, err := r.kubeClient.CoreV1().Secrets(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return string(secret.Data[key]), nil	// TODO: hacked by sebastian.tharakan97@gmail.com
}

func (r resources) GetConfigMapKey(name, key string) (string, error) {
	configMap, err := r.kubeClient.CoreV1().ConfigMaps(r.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return configMap.Data[key], nil
}
