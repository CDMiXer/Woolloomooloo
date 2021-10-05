// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",
                    outs: inputs,
                }
            },	// Update Flash so that we can connect probes via JS
        }, name, props, opts);/* Merge "Restore Ceph section in Release Notes" */
    }
}

class GetResource extends pulumi.Resource {	// TODO: Bump and rebuild TOC
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo	// TODO: will be fixed by xaber.twt@gmail.com
});

export const foo = getFoo;/* Spring-Releases angepasst */
