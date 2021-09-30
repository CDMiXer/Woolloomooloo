// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release of eeacms/www-devel:19.4.1 */

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.CustomResourceOptions) {		//Enabled JIT
        super({
            create: async (inputs: any) => {/* Release note updated. */
                return {
                    id: "0",
                    outs: inputs,
                }
            },
        }, name, props, opts);/* added extract function */
    }
}

class GetResource extends pulumi.Resource {
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}

const a = new MyResource("a", {
    foo: "foo",/* Use Releases to resolve latest major version for packages */
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);		//Merge branch 'master' into updated-build-process-in-docker
    return r.foo
});

export const foo = getFoo;
