import pulumi
import pulumi_kubernetes as kubernetes
/* Release 2.2.11 */
bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",/* Release OTX Server 3.7 */
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(/* Release 2.0.1 */
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
,)            
        )],	// Update leaflet/polygon-tabs/README.md
    ))
