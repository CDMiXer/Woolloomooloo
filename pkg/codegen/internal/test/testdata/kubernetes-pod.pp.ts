import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",	// add actionbar test and fixup stuff for gtk3
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {/* 477d4a4a-2e3a-11e5-8e6d-c03896053bdd */
                    memory: "20Mi",		//Remove RecyclerExceptionless
                    cpu: 0.2,
                },
            },
        }],
    },
});
