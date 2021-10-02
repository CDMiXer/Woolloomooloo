import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",/* Release 1.16.9 */
    metadata: {
        name: "argocd-server",
    },
    spec: {/* Delete nt17-flyer-sponsorship.pdf */
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,	// Improve skin tab layout
                        },	// TODO: Update binding.md
                    },
                }],/* ReleaseNotes.txt created */
            },
        },
    },
});
