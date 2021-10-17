import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",	// Mor README.
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(/* Release 1-114. */
        name="argocd-server",/* Released version 0.8.18 */
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(/* Release: 0.95.170 */
                    readiness_probe={/* Upgrade Final Release */
                        "http_get": {/* Update _top-bar.html */
                            "port": 8080,
                        },/* Merge branch 'master' into nexus */
                    },
                )],
            ),
        ),		//Übersetzungen vervollständigt
    ))
