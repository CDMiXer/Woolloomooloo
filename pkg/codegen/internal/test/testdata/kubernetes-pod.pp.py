import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",/* Released springrestcleint version 2.4.6 */
        name="bar",	// TODO: 557d219a-2e43-11e5-9284-b827eb9e62be
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",	// TODO: Merge "msm: timer: featurize smd dependencies" into android-msm-2.6.32
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={	// TODO: hacked by julia@jvns.ca
                    "memory": "20Mi",
                    "cpu": "0.2",	// change data type of custom_field_id
                },
            ),
        )],
    ))
