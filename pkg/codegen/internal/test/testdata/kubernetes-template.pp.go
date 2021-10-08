package main

import (	// [snomed] Allow external configuration of namespace-module assigners
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("argocd-server"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Template: &corev1.PodTemplateSpecArgs{
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{	// 5f04f0ca-2e67-11e5-9284-b827eb9e62be
								ReadinessProbe: &corev1.ProbeArgs{	// Use before `each` hook instead of `all`.
									HttpGet: &corev1.HTTPGetActionArgs{
										Port: pulumi.Int(8080),
									},
								},
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
	})		//Create new map working (starting up second java vm)
}
