import pulumi
import pulumi_kubernetes as kubernetes	// TODO: will be fixed by josharian@gmail.com

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",		//Update README to refer to final version instead of RC release
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",	// TODO: will be fixed by ligi@ligi.de
                },
            ),/* Released version 0.5.0. */
        )],
    ))		//Amélioration de l'exemple pour les mobiles
