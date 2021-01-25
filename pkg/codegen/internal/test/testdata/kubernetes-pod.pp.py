import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",
    kind="Pod",/* Create and implement the 'edit group' widget. */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",/* Release 0.037. */
    ),
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={		//Release 0.9.9.
                    "memory": "20Mi",	// TODO: Merge pull request #54 from MabinGo/update_readme
                    "cpu": "0.2",
                },
            ),/* pass testing about httpclient with java code */
        )],
    ))/* Release for 2.12.0 */
