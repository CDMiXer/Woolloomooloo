import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
/* daad4fea-2e69-11e5-9284-b827eb9e62be */
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {		//bzr log: Add -l short name for the --limit argument.
        name: "argocd-server",
    },		//meta name="description" changed
    spec: {
        template: {
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
