using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{/* Release v1.13.8 */
    public MyStack()
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs	// TODO: hacked by remco@dutchcoders.io
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",		//provisioning: add cost information
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {	// TODO: Just a few helper functions.
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs	// [XCore] Whitespace fixes, no functionality change.
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,		//Merge "[INTERNAL] sap.m.MultiInput: Removed odd class"
                                    },	// Use FaradayMiddleware::Mashify
                                },
                            },
                        },
                    },
                },
            },
        });
    }

}
