import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";
		//[ExoBundle] Modified table header
const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",/* Uploaded EM lecture */
    },
    spec: {/* Added ways to go from Easing to Gesture and from Gesture to Easing. */
        containers: [{	// Edit relations documents
            name: "nginx",
            image: "nginx:1.14-alpine",	// Add initial juju-chaos-monkey charm.
            resources: {
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],
    },
});
