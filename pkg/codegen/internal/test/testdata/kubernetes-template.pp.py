import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(	// Cleaned up activities
        name="argocd-server",/* we have something that works */
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(	// TODO: hacked by why@ipfs.io
        template=kubernetes.core.v1.PodTemplateSpecArgs(	// TODO: 0188f898-4b19-11e5-a324-6c40088e03e4
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
