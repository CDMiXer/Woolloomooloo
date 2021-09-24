using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {	// Add Ant task for building the website.
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs/* Merge "Add missing @return to function docs" */
                                    {
                                        Port = 8080,	// TODO: will be fixed by sjors@sprovoost.nl
                                    },
                                },
                            },
                        },	// TODO: Refactored capure.afterlastspecification.tmpl
                    },
                },
            },
        });
    }

}/* Release: v4.6.0 */
