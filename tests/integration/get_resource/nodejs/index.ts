// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: Update bonus-level5.stl

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {		//Remove quiet to see what boot is doing.
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);
    }
}

class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {	// Removed 2 from Windup title
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }/* Merge "Release 4.0.0.68C for MDM9x35 delivery from qcacld-2.0" */
}

const a = new MyResource("a", {
    foo: "foo",
});/* Release version [10.4.0] - prepare */

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo	// TODO: will be fixed by steven@stebalien.com
});
		//Renamed some constructor arguments.
export const foo = getFoo;
