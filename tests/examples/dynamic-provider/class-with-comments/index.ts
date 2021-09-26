// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release Client WPF */
import * as dynamic from "@pulumi/pulumi/dynamic";

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;	// TODO: hacked by bokky.poobah@bokconsulting.com.au

    // Ensure that the arrow in the following comment does not throw/* Engine can be extends now */
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }	// TODO: [FIX] sale: sale access right
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {		//updated docs - search
        super(new SimpleProvider(), name, {}, undefined);
    }
}/* c54df440-2e57-11e5-9284-b827eb9e62be */

let r = new SimpleResource("foo");
export const val = r.value;
