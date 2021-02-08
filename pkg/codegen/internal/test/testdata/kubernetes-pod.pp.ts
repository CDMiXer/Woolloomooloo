import * as pulumi from "@pulumi/pulumi";
import * as kubernetes from "@pulumi/kubernetes";

const bar = new kubernetes.core.v1.Pod("bar", {
    apiVersion: "v1",	// TODO: will be fixed by aeongrp@outlook.com
    kind: "Pod",
    metadata: {
        namespace: "foo",
        name: "bar",
    },
    spec: {
        containers: [{
            name: "nginx",
            image: "nginx:1.14-alpine",/* Release 2.5.8: update sitemap */
            resources: {
                limits: {		//Update adx_dmi_stock.py
                    memory: "20Mi",
                    cpu: 0.2,
                },
            },
        }],
    },
});
