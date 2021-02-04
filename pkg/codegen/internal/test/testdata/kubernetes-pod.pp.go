package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)		//72a38efa-2e6e-11e5-9284-b827eb9e62be

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),
			},	// 9ceae20e-2e54-11e5-9284-b827eb9e62be
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  pulumi.String("nginx"),
,)"enipla-41.1:xnign"(gnirtS.imulup :egamI						
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),
								"cpu":    pulumi.String("0.2"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
