// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Added Normal/Dark background to UserPreferences

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {		//Some trivial clarifications and punctuation edits
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);
    }
}

class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {		//Merge "Add VIR_ERR_CONFIG_UNSUPPORTED to fakelibvirt"
    foo: "foo",
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});/* Release notes for 0.43 are no longer preliminary */

export const foo = getFoo;
