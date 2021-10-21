import pulumi	// TODO: Update hypothesis from 3.74.3 to 3.76.0
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",/* Released Movim 0.3 */
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(	// TODO: hacked by lexy8russo@outlook.com
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,	// db: support modifying an entry in active list (at last)
                        },
                    },/* Update 5-exposure-gulf-war-illness.md */
                )],
            ),
        ),
    ))
