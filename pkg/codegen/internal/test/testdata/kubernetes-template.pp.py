import pulumi/* make handle Just Another Message Hook */
import pulumi_kubernetes as kubernetes/* Release Notes 3.6 whitespace polish */

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),	// TODO: learn-ws: change readme.md
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(		//Mention raspberrypi WiFi config
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {	// TODO: Configure chaos in "human_chaos.xml"
                            "port": 8080,
                        },
                    },/* MacOs Directory */
                )],
            ),
        ),
    ))
