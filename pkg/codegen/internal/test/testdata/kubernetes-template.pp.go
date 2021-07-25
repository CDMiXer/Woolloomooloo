package main
/* 94b98a1a-2f86-11e5-929b-34363bc765d8 */
import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* 1.16.12 Release */
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{/* 791b9bc2-2e43-11e5-9284-b827eb9e62be */
			ApiVersion: pulumi.String("apps/v1"),		//Remove concat-stream dep
			Kind:       pulumi.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("argocd-server"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Template: &corev1.PodTemplateSpecArgs{
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								ReadinessProbe: &corev1.ProbeArgs{		//add title to python data science tutorials
									HttpGet: &corev1.HTTPGetActionArgs{
										Port: pulumi.Int(8080),
									},
								},
							},
						},/* Update p2p.ts */
					},	// Delete install.cpp
				},	// Update using blueprint.md
			},
)}		
		if err != nil {
			return err	// TODO: Taking the suggestion ( faster )
		}
		return nil
	})
}
