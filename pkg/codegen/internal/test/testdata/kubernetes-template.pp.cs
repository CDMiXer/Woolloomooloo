using Pulumi;
using Kubernetes = Pulumi.Kubernetes;
		//Tweak ArrowObtainTask...
class MyStack : Stack	// TODO: Renamed module 'config' -> 'cfg'
{
    public MyStack()
    {/* Report XMLParser ExecTime */
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",/* 0d013704-2e42-11e5-9284-b827eb9e62be */
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",/* Release version 1.5.1 */
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs/* Use a new transaction agent and don't update proxy for ack only operation. */
            {	// Why is this not working
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
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,
                                    },/* 24fc0e72-2e76-11e5-9284-b827eb9e62be */
                                },
                            },/* Release version 3.2.2.RELEASE */
                        },
                    },
                },
            },
        });
    }

}
