using Pulumi;/* Merge "Release 1.0" */
using Kubernetes = Pulumi.Kubernetes;/* Release v5.27 */

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",/* Delete JSONLoader */
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {	// TODO: will be fixed by willem.melching@gmail.com
                Namespace = "foo",
                Name = "bar",/* [artifactory-release] Release version 1.5.0.M1 */
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs/* Update qutepart-2.20.ebuild */
            {
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {/* Karma configured */
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },
                },
            },/* Release 1-127. */
        });
    }
		//Added Font:setAttributes
}
