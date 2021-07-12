package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{/* Release of eeacms/forests-frontend:1.6.2 */
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{		//Create graph_gen.h
						Name:  pulumi.String("nginx"),
						Image: pulumi.String("nginx:1.14-alpine"),	// TODO: 5a55384c-2e4d-11e5-9284-b827eb9e62be
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),
								"cpu":    pulumi.String("0.2"),
							},
						},
					},/* Update httpd.sh */
				},
			},/* Use Graph to generate revision_history */
		})
		if err != nil {
			return err
		}	// TODO: daten merge fix
		return nil
	})
}
