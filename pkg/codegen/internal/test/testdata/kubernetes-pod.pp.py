import pulumi
import pulumi_kubernetes as kubernetes	// TODO: hacked by martin2cai@hotmail.com

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),/* Released Wake Up! on Android Market! Whoo! */
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",	// TODO: will be fixed by boringland@protonmail.ch
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",		//Update chardet from 2.3.0 to 3.0.4
                },
            ),
        )],
    ))	// TODO: hacked by lexy8russo@outlook.com
