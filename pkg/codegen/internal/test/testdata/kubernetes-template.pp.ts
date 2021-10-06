import * as pulumi from "@pulumi/pulumi";		//d942bffe-2e56-11e5-9284-b827eb9e62be
import * as kubernetes from "@pulumi/kubernetes";/* fixes layout test */
/* Release v0.22. */
const argocd_serverDeployment = new kubernetes.apps.v1.Deployment("argocd_serverDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",	// TODO: Update MogiiBot3.cs
    metadata: {
        name: "argocd-server",
    },/* display dataset columns information and allow to delete all columns */
    spec: {
        template: {
            spec: {
                containers: [{
                    readinessProbe: {/* UAF-3988 - Updating dependency versions for Release 26 */
                        httpGet: {
                            port: 8080,
                        },/* serialize/deserialize code moved to nblr-editor project */
                    },
                }],
            },	// Fix bug in TrackMatchingTest
        },
    },
});/* Merge "Support Library 18.1 Release Notes" into jb-mr2-ub-dev */
