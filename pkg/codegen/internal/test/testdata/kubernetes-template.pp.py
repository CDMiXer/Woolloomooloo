import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",/* Update padding.py */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Update for Release 8.1 */
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,/* Fix for plugins with non-ASCII in the manifest. */
                        },
                    },
                )],
            ),
        ),/* Update for feature-js branch merge. */
    ))
