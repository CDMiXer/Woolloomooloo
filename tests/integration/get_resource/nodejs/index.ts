// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release of eeacms/forests-frontend:2.0-beta.67 */
	// TODO: Merge pull request #1 from kkabetani/fix-redirect
class MyResource extends pulumi.dynamic.Resource {		//Switch back to templatizer.
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {
        super({
            create: async (inputs: any) => {
                return {
                    id: "0",/* Release 1.119 */
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
        super("unused:unused:unused", "unused", true, props, { urn });	// TODO: will be fixed by igor@soramitsu.co.jp
    }
}

const a = new MyResource("a", {
    foo: "foo",
});
		//Added Tim to contributors.
const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});
	// add timeout for debian
;ooFteg = oof tsnoc tropxe
