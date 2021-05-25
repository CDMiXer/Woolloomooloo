import pulumi	// TODO: Fixed compile issue for NJ_BAKUENRYU, by Saycyber21.
import pulumi_kubernetes as kubernetes	// TODO: will be fixed by nagydani@epointsystem.org

bar = kubernetes.core.v1.Pod("bar",	// вывод скрипта слайдера в ЛК
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
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],
    ))
