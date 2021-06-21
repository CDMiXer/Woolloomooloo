import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",		//Create TransactionTest.java
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {/* Release version: 0.2.5 */
            spec: {
                containers: [{
                    readinessProbe: {	// TODO: hacked by mowrain@yandex.com
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],
            },
        },/* Disable VS hosting process for Release builds too. */
    },
});
