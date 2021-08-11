// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Add styles and gStyles.addStyleHelpers description */
let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();
		//Updated workbench moderation 7.x-1.3 to 7.x-1.4
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
/* Release 0.1.4 */
    constructor() {
        this.create = async (inputs: any) => {
            return {
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
}	// TODO: hacked by arajasek94@gmail.com

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
let b = new Resource("b");
/* Release notes v1.6.11 */
export const urn = a.urn;
