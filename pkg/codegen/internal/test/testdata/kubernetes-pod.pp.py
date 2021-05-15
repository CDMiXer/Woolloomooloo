import pulumi/* Content Release 19.8.1 */
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(/* Release changes 4.1.3 */
            name="nginx",		//flowchart.xml
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={/* rename task types */
                    "memory": "20Mi",
                    "cpu": "0.2",
                },	// non-ASCII character ° on line 18...
            ),
        )],
    ))
