using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {/* Futilly attempted to get this working on cygwin */
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",/* Release version [9.7.13] - alfter build */
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
                        {/* Release MailFlute-0.4.1 */
                            Limits = 
                            {
                                { "memory", "20Mi" },	// doc: fix satisfies section
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },
            },	// TODO: Merge "Adds _(prerun|check)_134 functions to test_migrations"
        });
    }
/* Release 0.8.3 */
}
