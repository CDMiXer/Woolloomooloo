import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),	// Cleanup build.xml.
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(/* MessageRequestBuilder completed */
                limits={
                    "memory": "20Mi",/* Release of eeacms/www-devel:18.2.20 */
                    "cpu": "0.2",
                },
            ),
        )],
    ))
