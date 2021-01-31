// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release 0.13.1 (#703) */
import * as pulumi from "@pulumi/pulumi";/* Release v0.4.1 */

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,	// Merge "Run the v3 only job in neutron"
            };
        };
    }/* Preparing WIP-Release v0.1.39.1-alpha */
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
