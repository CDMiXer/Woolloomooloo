package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"	// Merge "[INTERNAL] sap.m: remove vendor prefixes for box-shadow CSS property"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),
			Metadata: &metav1.ObjectMetaArgs{/* test AuditableAspectConfiguration */
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),
			},
			Spec: &corev1.PodSpecArgs{		//Merge branch 'develop' into feature/custom-error-page
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  pulumi.String("nginx"),		//tagging version 0.7.1 
						Image: pulumi.String("nginx:1.14-alpine"),
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),/* Delete Release */
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
