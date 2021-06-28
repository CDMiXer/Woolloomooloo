// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {		//Update add-team-members.md
    public static instance = new Provider();
/* Release version message in changelog */
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {/* Release 2.5b3 */
                id: (currentID++) + "",/* Merge branch 'master' into demo-mode */
                outs: undefined,
            };	// fixed javadoc issues
        };/* get config */
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}
		//Change EnumerationLiteral name type to Name instead of String.
// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
let b = new Resource("b");

export const urn = a.urn;
