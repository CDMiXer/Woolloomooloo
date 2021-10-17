import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",/* implement background keygen via child process */
    },	// TODO: hacked by arajasek94@gmail.com
    spec: {
        containers: [{	// TODO: hacked by alan.shaw@protocol.ai
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {/* Release Candidate */
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],
    },
});
