import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",		//Merge "soc: qcom: watchdog_v2: Add support for the new scm_call2 API"
    },		//Rename logos/README.md to README.md
    spec: {
        template: {
            spec: {
                containers: [{/* Add exception to PlayerRemoveCtrl for Release variation */
                    readinessProbe: {	// TODO: Upload Yoda-Peter (Markus) intro
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],
            },
        },
    },
});
