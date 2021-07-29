import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",/* Release of eeacms/www:18.7.5 */
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
,}                        
                    },
                }],
            },
        },		//update photo urls for photobooth
    },
});
