// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* safeguard CB2 against talk page blanking */

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {	// Working sound implementation
                id: (currentID++) + "",
                outs: undefined,
            };
        };	// TODO: Fixed critical in substring call
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}/* [artifactory-release] Release version 1.2.0.RELEASE */

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
