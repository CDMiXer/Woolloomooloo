// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";/* Tagging a Release Candidate - v4.0.0-rc12. */

class SimpleProvider implements pulumi.dynamic.ResourceProvider {/* switched logo url to local */
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Release v1.0.6. */

    // Ensure that the arrow in the following comment does not throw
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {		//6d8bc32c-2e70-11e5-9284-b827eb9e62be
                id: "0",	// TODO: hacked by mikeal.rogers@gmail.com
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);		//Separated documentation of namespaces and whole project.
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
