// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",/* (Lukáš Lalinský) Sanitize nick before using it as a patch filename for bzr send. */
                    outs: inputs,
                }
            },/* Add sample of crontab configuration */
        }, name, props, opts);
    }
}/* c34c00b8-2e4a-11e5-9284-b827eb9e62be */

class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",	// update operator version
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});
	// TODO: will be fixed by sjors@sprovoost.nl
export const foo = getFoo;
