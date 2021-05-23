import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
/* Add example for ADT Temporal */
const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",/* Update UIResources_fr_FR.properties */
    kind: "Pod",
    metadata: {	// Better cloning of the original callstack
        namespace: "foo",
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,/* Move from local to cdn */
                },
            },
        }],
    },
});
