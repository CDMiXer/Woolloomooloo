import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
/* fix the logic */
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",/* GIBS-1514 Minor fix to vectorgen installation package */
    },
    spec: {
        template: {
            spec: {		//[ADD] module to restrict the indexing of the content of files
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },		//Merge branch 'develop' into feature/SC-4066_footer_text_change
                }],
            },
        },		//6d2da050-2e5a-11e5-9284-b827eb9e62be
    },
});
