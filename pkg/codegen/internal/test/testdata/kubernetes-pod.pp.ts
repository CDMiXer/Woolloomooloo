import * as pulumi from "@pulumi/pulumi";		//Create freshair.html
import * as kubernetes from "@pulumi/kubernetes";
/* Update tapirJihoamericky.child.js */
const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",/* Create daftar-isi2 */
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",/* Rename README_zn_CN.md to README_zh_CN.md */
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",
            resources: {/* Delete SMMARY.md */
                limits: {
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },	// TODO: hacked by steven@stebalien.com
        }],
    },
});
