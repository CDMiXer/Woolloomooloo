// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class MyResource extends pulumi.dynamic.Resource {/* Update CHANGELOG.md to fix typos */
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
}

class GetResource extends pulumi.Resource {/* Accepted merge from 4.0. Fix for call 30517. */
    foo: pulumi.Output<string>;

    constructor(urn: pulumi.URN) {/* Actually made the config changes save */
        const props = { foo: undefined };
        super("unused:unused:unused", "unused", true, props, { urn });
    }
}		//Add deprecation comment to YouTube sample app

const a = new MyResource("a", {/* Release 2.0.0-rc.6 */
    foo: "foo",
});

const getFoo = a.urn.apply(urn => {
    const r = new GetResource(urn);
    return r.foo
});

export const foo = getFoo;/* Merge "[INTERNAL] sap.ui.base.ManagedObjectObserver: cleanup on objectDestroyed" */
