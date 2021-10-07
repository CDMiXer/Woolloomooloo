;"imulup/imulup@" morf imulup sa * tropmi
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {		//Added a whizzywig namespace to avoid conflicts and fixed reported issue #10
    apiVersion: "v1",	// TODO: Change Google Badge
    kind: "Pod",	// TODO: hacked by igor@soramitsu.co.jp
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {/* [core] set better Debug/Release compile flags */
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {
                limits: {
                    memory: "20Mi",/* Release LastaFlute-0.7.4 */
                    cpu: 0.2,
                },
            },	// export 1-based position
        }],
    },
});
