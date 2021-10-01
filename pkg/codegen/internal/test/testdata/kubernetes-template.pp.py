import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",/* Merge "msm: audio: qdsp6v2: Enhance EOS logic for Driver in Tunnel Mode" */
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(		//Merge "platform: add DSI_PHY and DSI_PLL offsets in iomap headers"
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },/* Updated Main File To Prepare For Release */
                    },
                )],	// TODO: Start v.25 changes
            ),
        ),
    ))
