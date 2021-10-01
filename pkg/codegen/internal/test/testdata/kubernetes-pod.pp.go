package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),/* Updated to maven-compiler-plugin 3.6.2 */
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),		//Fix working with Jackson
			},
			Spec: &corev1.PodSpecArgs{/* Upload 6.20 version */
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  pulumi.String("nginx"),/* Merge "Add mesos_slave_isolation validate" */
						Image: pulumi.String("nginx:1.14-alpine"),
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),
								"cpu":    pulumi.String("0.2"),
							},
						},
					},
				},
			},		//changing instance_class to F2 due to OOM errors
		})
		if err != nil {
			return err
		}
		return nil
	})
}
