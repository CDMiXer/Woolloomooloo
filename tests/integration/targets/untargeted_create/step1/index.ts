// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release v1.4.6 */
import * as pulumi from "@pulumi/pulumi";/* update lang change */
/* Adhock Source Code Release */
let currentID = 0;
/* Merge "Add MFA Rules Release Note" */
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",/* js-core 2.8.1 RC1 released */
                outs: undefined,
            };
        };
    }		//Added and implemented LessThanOrEqualToOperator.
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
let b = new Resource("b");

export const urn = a.urn;
