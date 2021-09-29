package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"		//add HA link to config-settings toc
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"		//[robocompdsl] Minnor fix in import of test_dsl_factory.
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("argocd-server"),
,}			
			Spec: &appsv1.DeploymentSpecArgs{
				Template: &corev1.PodTemplateSpecArgs{
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								ReadinessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{/* Release of eeacms/www:20.6.26 */
										Port: pulumi.Int(8080),
									},
								},
							},	// Removing 0.9 branch.
						},/* 0513ff84-2e49-11e5-9284-b827eb9e62be */
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
