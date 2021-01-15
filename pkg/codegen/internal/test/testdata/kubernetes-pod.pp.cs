using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack	// Update lucia-manzano-girlsintech.html
{
    public MyStack()/* Delete mystery-aton.html */
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },/* minimOS-63 ABI stub */
                            },
                        },/* Release 1-78. */
                    },	// TODO: Moved all tests into a central folder. Moved public ct list to settings.
                },
            },
        });
    }

}
