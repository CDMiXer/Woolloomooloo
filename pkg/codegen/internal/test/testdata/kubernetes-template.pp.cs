using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{	// TODO: Update document index view to use updated_at
    public MyStack()/* add print_text_in_color */
    {
        var argocd_serverDeployment = new Kubernetes.Apps.V1.Deployment("argocd_serverDeployment", new Kubernetes.Types.Inputs.Apps.V1.DeploymentArgs
        {
            ApiVersion = "apps/v1",
            Kind = "Deployment",	// TODO: hacked by alex.gaynor@gmail.com
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Name = "argocd-server",	// TODO: hacked by witek@enjin.io
            },	// TODO: will be fixed by ac0dem0nk3y@gmail.com
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
                                {	// TODO: Merge branch 'master' into remove-warn
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
        });
    }

}
