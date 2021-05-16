using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",
            },/* Log to MumbleBetaLog.txt file for BetaReleases. */
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {/* Update README Release History */
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {/* Create ISB-CGCBigQueryTableSearchReleaseNotes.rst */
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {/* Release for 19.1.0 */
                                        Port = 8080,
                                    },
                                },/* Add unit tests about sent mails. */
                            },
                        },
                    },
                },
            },
        });
    }

}
