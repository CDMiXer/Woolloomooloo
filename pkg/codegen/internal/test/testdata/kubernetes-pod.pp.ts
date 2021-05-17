import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",/* Merge "[INTERNAL] Release notes for version 1.30.5" */
    kind: "Pod",		//TICTOCK - DEBUG
    metadata: {
        namespace: "foo",/* Defining types */
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],/* gradient background */
    },
});
