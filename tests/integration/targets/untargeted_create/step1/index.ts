// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Fix busted docs. */
import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();	// TODO: hacked by witek@enjin.io
	// TODO: will be fixed by indexxuan@gmail.com
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;		//Enable advanced AI in SP

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,		//Fix windows paths in TsParser
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
let b = new Resource("b");	// Pull the planners out of install and pass them as parameters
	// Applied to trunk - Fix to new species for site verification
export const urn = a.urn;	// TODO: testnet genesis block
