import pulumi
import pulumi_kubernetes as kubernetes		//Improve readme.

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(		//[ELASTICMS-39] add entity notification
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(		//remove name field from Binding
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={
                        "http_get": {
                            "port": 8080,/* Environment for simple graph search */
                        },
                    },
                )],/* Merge branch 'master' into fix-links */
            ),
        ),		//clarify information; add a link to the Places
    ))
