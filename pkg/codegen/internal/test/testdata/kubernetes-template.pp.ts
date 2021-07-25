import * as pulumi from "@pulumi/pulumi";/* Upgrade version number to 3.1.4 Release Candidate 1 */
import * as kubernetes from "@pulumi/kubernetes";
	// Move code snippet into 'bash' section
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {	// Fix that the FormDescriptions decorator does not ignore Notes and Buttons
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",
    },
    spec: {/* Merge "Release 3.2.3.333 Prima WLAN Driver" */
        template: {
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],		//WwBM3F75vbuQQJsJCNaxRJ8mWcS7O9gN
            },
        },
    },
});
