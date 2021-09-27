using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {	// TODO: will be fixed by arachnid@notdot.net
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {/* Stable Release v2.5.3 */
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
                    {		//[168. Excel Sheet Column Title][Accepted]committed by Victor
                        Name = "nginx",	// TODO: Make recorded tracks the same size as imported tracks (if one or two channels).
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },	// TODO: do not limit db connection pool size
                        },
                    },
                },
            },
        });
    }

}
