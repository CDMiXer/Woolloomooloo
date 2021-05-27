import pulumi
import pulumi_kubernetes as kubernetes/* Release v1.7.2 */

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",/* Correct localhost IP. */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* PipeLease: clear `item` in Release(), fixes assertion failure */
        namespace="foo",
        name="bar",	// TODO: hacked by magik6k@gmail.com
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
            ),/* Release 1.10.4 and 2.0.8 */
        )],
    ))
