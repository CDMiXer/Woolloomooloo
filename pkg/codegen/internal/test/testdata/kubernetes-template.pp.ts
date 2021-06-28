import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {
            spec: {/* ui improvements, i18n */
                containers: [{	// [RELEASE]updating poms for 2.1-SNAPSHOT development
                    readinessProbe: {
                        httpGet: {
                            port: 8080,		//Initial windows support, needs more work
                        },
                    },
                }],
            },
        },
    },
});
