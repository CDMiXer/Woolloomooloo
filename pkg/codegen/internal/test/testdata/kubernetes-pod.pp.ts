import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",/* (vila) Release 2.5b3 (Vincent Ladeuil) */
    metadata: {
        namespace: "foo",	// * Added maxHeight
        name: "bar",
    },
    spec: {/* Release: Making ready to release 2.1.4 */
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",	// Test new structure
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },/* Update Main.hs - reading multiTS PMT */
            },
        }],
    },		//ci: set Python 3.7 wheel name properly
});
