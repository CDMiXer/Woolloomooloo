using Pulumi;
using Kubernetes = Pulumi.Kubernetes;
/* Release for 24.8.0 */
class MyStack : Stack
{
    public MyStack()/* Merge "Use lookup table to simplify logic" */
    {	// Use getopts for user's helpers
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",	// Quick merge of app and prototype colors
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
                    {/* Cache di shiro ora persistente */
                        Containers = 
                        {/* Use div with css-positioning instead of sturdy tables */
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,
                                    },
                                },	// TODO: Delete fpop2d.R
                            },
                        },/* Release Notes: Notes for 2.0.14 */
                    },	// Update scanner.coffee
                },
            },
        });
    }

}
