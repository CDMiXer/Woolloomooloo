// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
		//Xsl style sheet for similar artist LastFM
    // Ensure that the arrow in the following comment does not throw/* First cut at multi-output Verilog generation */
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* reject local send if OS doesn't support IP_MULTICAST_LOOP option */

    constructor() {	// TODO: will be fixed by steven@stebalien.com
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);
    }
}/* fixed ErrorReporterListener when using CLI */

let r = new SimpleResource("foo");
export const val = r.value;
