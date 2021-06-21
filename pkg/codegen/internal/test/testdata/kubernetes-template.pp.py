import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(	// TODO: settings routing spec stub
        template=kubernetes.core.v1.PodTemplateSpecArgs(/* fix(package): update prismjs to version 1.13.0 */
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },
                    },	// TODO: Optionally cache volume to reduce number of volume get requests
                )],
            ),
        ),
))    
