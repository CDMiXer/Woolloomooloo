import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",	// TODO: Rename project/connectome_learn to Project/connectome_learn
    ),/* Release: version 1.1. */
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(		//Added services.json to .gitignore
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],
    ))/* Release of eeacms/eprtr-frontend:0.3-beta.14 */
