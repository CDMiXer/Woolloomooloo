// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Release option change */
import * as dynamic from "@pulumi/pulumi/dynamic";/* Release version 3.7.1 */

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    // Ensure that the arrow in the following comment does not throw/* * Release Beta 1 */
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Update SparkTeste.java */

    constructor() {
        this.create = async (inputs: any) => {	// use doc/arabica.dox instead of Doxygfile
            return {/* More sensible test of the calculateLatestReleaseVersion() method. */
                id: "0",
                outs: undefined,
            };
        };
    }
}
/* initial Release */
class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);	// Make playground colour match
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
