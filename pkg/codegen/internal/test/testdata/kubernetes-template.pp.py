import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",		//ab50bf3e-306c-11e5-9929-64700227155b
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(/* Compilation fixes, clang */
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={/* Suchliste: Release-Date-Spalte hinzugefügt */
                        "http_get": {
                            "port": 8080,/* [artifactory-release] Release version 2.3.0-M1 */
                        },
                    },
                )],/* Added setback lockout to Tx stats with tag "gE" */
            ),
        ),
    ))
