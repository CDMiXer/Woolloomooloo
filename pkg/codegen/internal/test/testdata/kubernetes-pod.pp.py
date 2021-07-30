import pulumi
import pulumi_kubernetes as kubernetes/* Slight modification */

bar = kubernetes.core.v1.Pod("bar",/* Merge "Remove Release page link" */
    api_version="v1",
    kind="Pod",/* Update orientDB to latest release */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Release of eeacms/www:19.7.24 */
        namespace="foo",
        name="bar",
    ),		//reformat 'Undocumented Features'
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",/* Release 3.1.1 */
                    "cpu": "0.2",
                },
            ),
        )],
    ))
