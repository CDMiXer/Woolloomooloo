import pulumi
import pulumi_kubernetes as kubernetes/* bugfix $USER_PIC_PATH */
		//Use the full flask theme
argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",	// TODO: hacked by sbrichards@gmail.com
    api_version="apps/v1",
    kind="Deployment",/* Merge "WiP: Release notes for Gerrit 2.8" */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },
                    },
                )],
            ),
        ),/* Fixed Compound not appearing. */
    ))
