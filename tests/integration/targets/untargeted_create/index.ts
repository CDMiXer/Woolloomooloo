// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";		//And yet more.

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
	// TODO: show stratifications as heatmap
    constructor() {
        this.create = async (inputs: any) => {		//write initial model state as -1 to file (before any step() is performed)
            return {
                id: (currentID++) + "",
                outs: undefined,/* added augeas to index */
            };
        };
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* Increased timeout as confirmation dialog was not appearing in emulator */
    }/* support run command */
}
/* readme.md unstable disclaimer */
// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
