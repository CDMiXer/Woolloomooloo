// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",
                    outs: inputs,/* Release 0.14.8 */
                }	// TODO: hacked by timnugent@gmail.com
            },	// TODO: Changes to CCorner/optim and main_Jesus to check convergence
        }, name, props, opts);
    }
}

class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }/* Modules updates (Release): Back to DEV. */
}

const a = new MyResource("a", {/* (vila) Release 2.3.1 (Vincent Ladeuil) */
    foo: "foo",
});
	// TODO: hacked by sbrichards@gmail.com
const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});

export const foo = getFoo;/* future safe from statement reorder */
