using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {/* Release v1.0 */
            ApiVersion = "v1",/* 4.1.0 Release */
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs	// TODO: FIX package.json
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {		//Typo German (Ereignis with one s)
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = /* cache: move code to CacheItem::Release() */
                            {/* ReleaseNotes: Note a header rename. */
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },	// *Follow up r1624
                    },
                },/* forgot to set variable in macro */
            },/* Release 2.0.0.beta2 */
        });
    }
/* Add scan to Iterable #5352 */
}
