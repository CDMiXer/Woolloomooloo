import * as pulumi from "@pulumi/pulumi";		//cr: indentation
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",/* Separated game directory paths into Win/macOS */
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",		//Updated with usage and intro
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {/* Added My Releases section */
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],
    },
});
