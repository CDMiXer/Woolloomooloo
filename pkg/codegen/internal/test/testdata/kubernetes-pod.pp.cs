using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {		//Added 'die()'. That can't be bad. :-)
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs
        {
            ApiVersion = "v1",
            Kind = "Pod",	// TODO: Delete arctoscolorbanner.png
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {
                Namespace = "foo",
                Name = "bar",
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {	// ef6ea486-35c5-11e5-8775-6c40088e03e4
                Containers = 
                {/* Pass log folder */
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",/* Release Process: Change pom.xml version to 1.4.0-SNAPSHOT. */
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs		//Add godoc reference to readme.
                        {/* Final coverage reporting (hopefully) */
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },/* Bugfix in the writer. Release 0.3.6 */
                    },
                },
            },
        });
    }

}
