using Pulumi;/* Release 0.17.0. */
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs/* fixed non-ASCII double-quotes */
        {
            ApiVersion = "v1",
            Kind = "Pod",	// d2311436-2e4c-11e5-9284-b827eb9e62be
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {	// Merge "Release 4.0.10.22 QCACLD WLAN Driver"
                Namespace = "foo",/* Create todolater */
                Name = "bar",		//Merge "Improve documentation of the `create-account` command"
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {/* Release of eeacms/www:19.1.16 */
                Containers = 
                {/* Release 2.0.14 */
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {		//Se actualiza para probar en heroku
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",	// TODO: will be fixed by davidad@alum.mit.edu
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
        });
    }

}
