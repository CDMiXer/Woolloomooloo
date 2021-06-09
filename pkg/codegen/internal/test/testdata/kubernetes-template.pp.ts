import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",/* Merge "Release notes: deprecate kubernetes" */
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {/* Merge "Choose a server group when booting a VM with NG launch instance" */
            spec: {	// fix lib_dir sed typo
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],
            },
        },
    },
});
