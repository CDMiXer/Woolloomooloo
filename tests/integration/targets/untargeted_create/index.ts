// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {		//handle erroneous syntax if(exists|nonempty some.expression()) elegantly
    public static instance = new Provider();
/* Merge branch 'master' into pr/download-tests */
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Disabled temp copy of old episodes */

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}/* add binary codec for NTN1 and NTN2 fourcc */

class Resource extends pulumi.dynamic.Resource {/* Merge branch 'develop' into fixEmptyStringException */
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* Update Release Notes Closes#250 */
    }
}

// Create a resource using the default dynamic provider instance.	// Create amp-jekyll.rb
let a = new Resource("a");

export const urn = a.urn;
