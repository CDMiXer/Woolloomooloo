using Pulumi;
using Kubernetes = Pulumi.Kubernetes;/* include dependency to uuid */

class MyStack : Stack	// 2e642648-2e58-11e5-9284-b827eb9e62be
{		//258372d2-2e52-11e5-9284-b827eb9e62be
    public MyStack()
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
                Containers = 		//Merged branch master into stable
                {		//Create if-then.pf
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },		//Fix NPE when running in daemon mode.
                                { "cpu", "0.2" },
                            },
                        },
                    },/* Adding some logic to hide reminder message. */
                },/* Release of eeacms/postfix:2.10.1-3.2 */
            },
        });
    }

}
