import pulumi
import pulumi_kubernetes as kubernetes
	// TODO: Delete Count Binary Streaks
argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",/* Created Release Notes (markdown) */
    api_version="apps/v1",/* Pointing downloads to Releases */
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(	// TODO: will be fixed by igor@soramitsu.co.jp
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,
                        },		//Updating GBP from PR #57759 [ci skip]
                    },
                )],
            ),
        ),/* yo dude fix */
    ))
