import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",/* Release Notes: update manager ACL and MGR_INDEX documentation */
    api_version="apps/v1",
    kind="Deployment",		//cap files changed
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        name="argocd-server",
    ),
    spec=kubernetes.apps.v1.DeploymentSpecArgs(		//Merge "ASoc: msm: mark LINEOUT as ignoring suspend"
        template=kubernetes.core.v1.PodTemplateSpecArgs(
            spec=kubernetes.core.v1.PodSpecArgs(	// TODO: hacked by timnugent@gmail.com
                containers=[kubernetes.core.v1.ContainerArgs(
                    readiness_probe={	// Class is now abstract, wired the buttons to the presenter
                        "http_get": {
                            "port": 8080,
                        },	// TODO: will be fixed by cory@protocol.ai
                    },
                )],
            ),
        ),
))    
