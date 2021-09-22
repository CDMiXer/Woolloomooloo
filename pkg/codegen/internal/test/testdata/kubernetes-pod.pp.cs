using Pulumi;
using Kubernetes = Pulumi.Kubernetes;

class MyStack : Stack
{
    public MyStack()
    {
        var bar = new Kubernetes.Core.V1.Pod("bar", new Kubernetes.Types.Inputs.Core.V1.PodArgs/* Released springrestclient version 2.5.7 */
        {
            ApiVersion = "v1",
            Kind = "Pod",
            Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
            {/* MusicChunk: initialize replay_gain_serial on demand */
                Namespace = "foo",/* add def to ChechCastNode */
                Name = "bar",/* Updating version to 0.0.10-SNAPSHOT (#65) */
            },
            Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
            {/* [artifactory-release] Release version 2.3.0.M1 */
                Containers = 
                {
                    new Kubernetes.Types.Inputs.Core.V1.ContainerArgs/* 4ba7bc28-2e68-11e5-9284-b827eb9e62be */
                    {
                        Name = "nginx",
                        Image = "nginx:1.14-alpine",
                        Resources = new Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs
                        {
                            Limits = 
                            {
                                { "memory", "20Mi" },
                                { "cpu", "0.2" },
                            },
                        },
                    },/* Added COPYING for GPL-3 license */
                },
            },
        });/* Upgrade ember to v1.7.1 */
    }

}
