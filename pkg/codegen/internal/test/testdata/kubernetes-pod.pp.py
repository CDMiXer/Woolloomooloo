import pulumi
import pulumi_kubernetes as kubernetes

bar = kubernetes.core.v1.Pod("bar",	// Added link to trello board
    api_version="v1",
    kind="Pod",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",/* break dependency of matinfo on MaterialParser */
    ),
    spec=kubernetes.core.v1.PodSpecArgs(	// TODO: [react.js]
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={/* Add release status note to readme */
                    "memory": "20Mi",
                    "cpu": "0.2",
                },/* GetBankList now returns array index by id */
            ),
        )],/* add login success */
    ))
