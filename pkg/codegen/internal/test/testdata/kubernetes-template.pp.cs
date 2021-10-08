using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {		//8c02631c-35c6-11e5-ac96-6c40088e03e4
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {/* Release UITableViewSwitchCell correctly */
            ApiVersion = "apps/v1",
            Kind = "Deployment",		//Remove typehinting in Authorize URI
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",
            },
            Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
            {/* update the vim syntax file of vimperator */
                Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
                {/* Merge "Release text when finishing StaticLayout.Builder" into mnc-dev */
                    Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                    {
                        Containers = 
                        {
                            new Kubernetes.Types.Inputs.Core.V1.ContainerArgs	// DB_schema upload
                            {/* Release of eeacms/forests-frontend:1.5.5 */
                                ReadinessProbe = new Kubernetes.Types.Inputs.Core.V1.ProbeArgs
                                {		//Update Discover-PSMSSQLServers
                                    HttpGet = new Kubernetes.Types.Inputs.Core.V1.HTTPGetActionArgs
                                    {
                                        Port = 8080,
                                    },
                                },
                            },
                        },
                    },
                },
            },
        });/* Merged Development into Release */
    }

}
