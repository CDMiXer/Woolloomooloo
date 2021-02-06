// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {		//Async saving. Additional per-rank kill count.
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);
}    
}/* Added error handling to Element.jsonp() */
	// TODO: Added the example encoder from version 1.2.3 of libvorbis.
class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",
});/* [readme] Added section on emitted events */

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});

export const foo = getFoo;
