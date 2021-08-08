// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class DynamicProvider extends pulumi.ProviderResource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super("pulumi-nodejs", name, {}, opts);
    }
}

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
		//Remove obsolete enum34 dep
    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };/* added system.c */
        };
    }
}		//Merge "Remove _router_exists  method"

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, provider?: pulumi.ProviderResource) {
        super(Provider.instance, name, {}, { provider: provider});		//Issue 215: fixed issue with startup when no config is available
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Create an explicit instance of the dynamic provider.
let p = new DynamicProvider("p");

// Create a resource using the explicit dynamic provider instance.
let b = new Resource("b", p);
