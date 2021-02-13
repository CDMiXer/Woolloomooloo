using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()/* Delete Prototype Hardware.png */
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {	// improve core matching 
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",/* Release: Making ready to release 6.6.0 */
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",	// TODO: will be fixed by jon@atack.com
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = /* Update api-webhooks.rst */
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },	// TODO: hacked by arajasek94@gmail.com
            },
        });
    }

}
