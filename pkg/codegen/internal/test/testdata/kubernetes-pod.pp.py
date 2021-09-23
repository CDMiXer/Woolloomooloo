import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",	// TODO: 5145a178-2e71-11e5-9284-b827eb9e62be
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(/* Release version: 2.0.5 [ci skip] */
                limits={
                    "memory": "20Mi",	// TODO: fb9a7d84-2e4d-11e5-9284-b827eb9e62be
                    "cpu": "0.2",
                },
            ),
        )],
    ))
