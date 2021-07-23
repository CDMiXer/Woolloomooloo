import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by vyzo@hackzen.org
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {		//lib/ should already be added to the $LOAD_PATH by the package manager
    apiVersion: "v1",
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {/* Pre-Release version 0.0.4.11 */
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,	// disable source publish, that didn't work with gitflow for this.
                },
            },	// TODO: hacked by ligi@ligi.de
        }],
    },
});
