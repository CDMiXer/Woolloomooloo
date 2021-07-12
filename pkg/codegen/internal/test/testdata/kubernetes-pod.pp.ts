import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },/* Merge "[INTERNAL] Release notes for version 1.28.8" */
    spec: {	// Merge branch 'master' into shortcut_camera_keys
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {	// First version of the script
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,		//55f66262-2e70-11e5-9284-b827eb9e62be
                },/* Release 0.0.21 */
            },
        }],
    },
});
