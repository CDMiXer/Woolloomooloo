package main

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: pulumi.String("apps/v1"),
			Kind:       pulumi.String("Deployment"),		//moved irix stuff above cpack, etc
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("argocd-server"),
			},/* Syntax coloring for C++ snippets in README.md */
			Spec: &appsv1.DeploymentSpecArgs{/* primo esercizio vettori */
				Template: &corev1.PodTemplateSpecArgs{	// TODO: hacked by igor@soramitsu.co.jp
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								ReadinessProbe: &corev1.ProbeArgs{		//New translations en-GB.mod_sermonupload.sys.ini (Spanish, Colombia)
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
		if err != nil {	// TODO: will be fixed by steven@stebalien.com
			return err
		}
		return nil
	})
}
