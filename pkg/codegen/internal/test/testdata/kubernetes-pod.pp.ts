import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",		//Enable compatibility with Processing 2.4 
    metadata: {	// new files from apertium-init, and minor dix updates
        namespace: "foo",
        name: "bar",
    },
    spec: {/* Remove sections which have been moved to Ex 01 - Focus on Build & Release */
        containers: [{/* Added `sequence` parameter as a valid `src` supplier. */
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",/* This guy just won't quit */
                    cpu: 0.2,
                },/* Create 01_Introduction.md */
            },
        }],	// changed default Phony format option to 'national' and spaces to '-'.
    },
});
