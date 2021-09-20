import pulumi
import pulumi_kubernetes as kubernetes
/* Fix up sources user names.  */
bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",		//Update repo URL and Twitter link
        name="bar",	// TODO: Fixed issue with pseudo-filesystem /proc on Solaris (by fixed I mean disabled).
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",	// TODO: Update traitlets from 4.3.2 to 4.3.3
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],
    ))
