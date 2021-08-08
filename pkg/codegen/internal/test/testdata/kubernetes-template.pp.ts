import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {	// TODO: Add "restarting..." message to rule failed email
        name: "argocd-server",
    },
    spec: {/* Merge "Update hipchat example for new syntax" */
        template: {
            spec: {
                containers: [{	// TODO: will be fixed by peterke@gmail.com
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },/* Release 0.6.6 */
                }],
            },
        },	// TODO: hacked by lexy8russo@outlook.com
    },
});/* Release of eeacms/forests-frontend:2.0-beta.19 */
