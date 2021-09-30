import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
/* Release 1.04 */
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "argocd-server",		//+страница авторизации
    },
    spec: {
        template: {
            spec: {	// trimmed log output
                containers: [{/* TextViewWithCharacterLimitLabelDelegate added */
                    readinessProbe: {/* Added schema.org information to the user profile. */
                        httpGet: {
                            port: 8080,
                        },
                    },		//Added functionality to search Google Scholar when pressing backslash
                }],
            },
        },
    },/* @Release [io7m-jcanephora-0.9.17] */
});/* Added isEmpty() */
