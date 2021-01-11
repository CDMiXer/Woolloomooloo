import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {/* Release of XWiki 12.10.3 */
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {
        template: {/* Fix travis : git-{rebase,commit}.el and with-editor.el are already in magit */
            spec: {
                containers: [{
                    readinessProbe: {		//Update README.md with 0.9.2 info
                        httpGet: {
                            port: 8080,/* Release script updated. */
                        },
                    },
                }],
            },
        },
    },
});
