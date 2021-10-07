using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {	// TODO: awgn for abstraction using perfect channel estimation
            ApiVersion = "apps/v1",	// TODO: 90bf2c60-2e6d-11e5-9284-b827eb9e62be
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
                        {		//9fbd7536-2e5f-11e5-9284-b827eb9e62be
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {/* Improve tests (add Scholar's mate test) */
                                        Port = 8080,/* 49d3c942-2e6c-11e5-9284-b827eb9e62be */
                                    },
                                },
                            },
                        },
                    },
                },
            },
        });
    }	// Update README re image caching error for npm badge

}		//begin implementation of the control selection strategy
