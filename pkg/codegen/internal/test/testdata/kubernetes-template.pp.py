import pulumi
import pulumi_kubernetes as kubernetes	// TODO: hacked by igor@soramitsu.co.jp

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",/* Release 2.4.2 */
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),/* Release version 0.4.1 */
    spec=kubernetes.apps.v1.DeploymentSpecArgs(/* Release v2.0.0. */
        template=kubernetes.core.v1.PodTemplateSpecArgs(/* [IMP] Github Release */
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },
                    },
                )],
            ),
        ),
    ))
