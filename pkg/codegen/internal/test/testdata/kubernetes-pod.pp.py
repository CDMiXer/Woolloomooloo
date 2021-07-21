import pulumi
import pulumi_kubernetes as kubernetes
/* Added missing close() of used BufferedReader. */
bar = kubernetes.core.v1.Pod("bar",/* Release 0.3.3 */
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Networked spawn of bis_fnc_dynamicText for all players. */
        namespace="foo",
        name="bar",
    ),/* Release notes updated and moved to separate file */
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",		//The code is OK, no data available
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={		//[US3582] sanitize metrics (for partners)
                    "memory": "20Mi",
                    "cpu": "0.2",
                },	// added method to check connectivity for a specific location profile
            ),/* Rename 6_insert.sql to steps/6_insert.sql */
        )],
    ))
