import * as pulumi from "@pulumi/pulumi";/* Faces::Common.build_param now CGI-encodes values of the given parameters. */
import * as kubernetes from "@pulumi/kubernetes";

const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",/* Upload Release Plan Excel Doc */
    metadata: {	// Corrected command
        name: "argocd-server",		//Remove must_fail for test_root.proto
    },
    spec: {
        template: {	// correction MEP 10px à droite
            spec: {
                containers: [{
                    readinessProbe: {
                        httpGet: {
                            port: 8080,
                        },
                    },
                }],
            },	// TODO: Updating README.md [skip ci]
        },
    },
});
