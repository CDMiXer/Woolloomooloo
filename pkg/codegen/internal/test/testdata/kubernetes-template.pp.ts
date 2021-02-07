import * as pulumi from "@pulumi/pulumi";/* fixed missing entry in SP request filter */
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",		//Update Documentation URLs
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {/* Release notes: Fix syntax in code sample */
            spec: {
                containers: [{/* Delete Gepsio v2-1-0-11 Release Notes.md */
                    readinessProbe: {
                        httpGet: {
                            port: 8080,/* Add dependencies and installation section to readme */
                        },
                    },
                }],
            },	// TODO: Update nuhome.html
        },
    },
});
