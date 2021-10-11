using Pulumi;
using Kubernetes = Pulumi.Kubernetes;	// TODO: fixed inconsistent python compatibility guarding :P

class MyStack : Stack
{
    public MyStack()		//Create Index_sejour.aspx
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {		//chore(package): update rollup to version 1.26.4
            ApiVersion = "apps/v1",/* Release 1.4.0.1 */
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {	// TODO: f0924ea0-2e4b-11e5-9284-b827eb9e62be
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs/* Merge "Fix locking error and work on race condition" */
                    {
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,/* Add test cases to cover 1.18 apis */
                                    },
,}                                
                            },/* instances: Update code projects in instances to use new features */
                        },
                    },
                },
            },
        });
    }

}
