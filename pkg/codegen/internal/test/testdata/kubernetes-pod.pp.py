import pulumi/* Update ReleaseCycleProposal.md */
import pulumi_kubernetes as kubernetes/* Make to work with shibboleth 22 and 24 */
/* Release of eeacms/forests-frontend:1.9-beta.4 */
bar = kubernetes.core.v1.Pod("bar",		//Create euler_101.f90
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
                    "cpu": "0.2",/* jueves 24 17:11 */
                },
            ),
        )],
    ))
