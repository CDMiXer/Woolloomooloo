import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";	// TODO: hacked by vyzo@hackzen.org

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {/* 7a0dc314-2e72-11e5-9284-b827eb9e62be */
        name: "argocd-server",
    },
    spec: {
        template: {		//fix for left shift of Word64
            spec: {
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
