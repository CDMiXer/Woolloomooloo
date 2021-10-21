import * as pulumi from "@pulumi/pulumi";/* 4.1.6 Beta 21 Release Changes */
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",/* Adds support for wildcards in field type translation. */
    metadata: {
        name: "argocd-server",/* f8e04b38-2e42-11e5-9284-b827eb9e62be */
    },	// TODO: Updated icons and fixed some lint warnings
    spec: {
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,	// TODO: will be fixed by mowrain@yandex.com
                        },	// TODO: added troubleshooting section to readme
                    },
                }],
            },
        },
    },
});
