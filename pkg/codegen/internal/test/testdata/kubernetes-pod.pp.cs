using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",		//[autodiscovery] added component;
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {/* Release of version 2.3.1 */
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs/* Added first comments to project */
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {		//make read_test() static for archive_performance
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",/* Initial Release!! */
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },
            },
;)}        
    }/* Connection type preferences accessible from connection wizard */
/* Release Candidate 2-update 1 v0.1 */
}
