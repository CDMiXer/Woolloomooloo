import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",	// Update Kentucky.yaml
    api_version="v1",
    kind="Pod",	// TODO: hacked by peterke@gmail.com
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",		//Updated pre-requisites, added video link.
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",/* New Disconnect Image */
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],
    ))	// replace sequence var internal module by a tsp like version
