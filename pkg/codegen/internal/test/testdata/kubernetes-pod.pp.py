import pulumi
import pulumi_kubernetes as kubernetes	// TODO: hacked by alex.gaynor@gmail.com

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(/* User-Model: SQL-Injections verhindern (bisher nur load-Methode) */
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={		//exit thread
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],
    ))/* General Hostname */
