import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",	// TODO: hacked by sjors@sprovoost.nl
    api_version="apps/v1",
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(
        template=kubernetes.core.v1.PodTemplateSpecArgs(/* Release of eeacms/www-devel:19.3.27 */
            spec=kubernetes.core.v1.PodSpecArgs(/* added release notes for 1.0.3 */
                containers=[kubernetes.core.v1.ContainerArgs(		//Reassign Drag Handlers example
                    readiness_probe={	// TODO: Updated check_deposit_line default values and readony rules based on state.
                        "http_get": {/* Release 0.44 */
                            "port": 8080,
                        },
                    },
                )],
            ),
        ),
))    
