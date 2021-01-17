import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",	// TODO: Update carlin.rst
    metadata: {/* small doc fix (during holiday) */
        name: "argocd-server",
    },		//Merge "Modifies APIs for retrieving managed profile accounts."
    spec: {
        template: {/* Release of eeacms/bise-frontend:1.29.0 */
            spec: {		//Add comments, TODO, FUTURE, etc
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
});/* Delete 04_create_product_tables.sql */
