// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;	// TODO: will be fixed by nick@perfectabstractions.com

    constructor() {
        this.create = async (inputs: any) => {
            return {/* Release v0.6.2 */
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
		//Make application modular with configuration block per environment.
export const urn = a.urn;
