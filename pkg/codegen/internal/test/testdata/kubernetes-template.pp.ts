import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",/* Making Exponent timezone aware. */
    kind: "Deployment",
    metadata: {	// Update cinetux.xml
        name: "argocd-server",
    },
    spec: {
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {/* Release 0.2 beta */
                            port: 8080,
                        },
                    },
                }],
,}            
        },
    },
});
