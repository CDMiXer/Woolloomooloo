import pulumi
import pulumi_kubernetes as kubernetes

argocd_server_deployment = kubernetes.apps.v1.Deployment("argocd_serverDeployment",	// TODO: hacked by brosner@gmail.com
    api_version="apps/v1",	// http://englishplus.com/grammar/00000296.htm
    kind="Deployment",
    metadata=kubernetes.meta.v1.ObjectMetaArgs(	// TODO: hacked by aeongrp@outlook.com
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
                )],		//Alteração das variaveis da tela clientes
            ),
        ),	// Add blackbox test for displaying non-ascii log in bzr version
    ))
