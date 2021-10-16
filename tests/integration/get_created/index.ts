// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
		//Added cran
    constructor() {
        this.create = async (inputs: any) => {
            return {	// TODO: version bump 0.2.0
                id: "0",
                outs: undefined,
            };
        };
    }
}/* Release: Making ready for next release cycle 5.0.1 */

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Attempt to read the created resource.	// TODO: will be fixed by yuvalalaluf@gmail.com
let b = new Resource("b", { id: a.id });
