// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* added demo plunker */

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();
	// TODO: hacked by mail@bitpshr.net
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {	// [doc] add some inline explanation of t2
        this.create = async (inputs: any) => {/* Release 1.0.28 */
            return {	// TODO: hacked by nicksavers@gmail.com
                id: "0",
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

// Create a resource using the default dynamic provider instance./* Ignore crawler log file */
let a = new Resource("a");
/* dr73: #i103152# description of T() */
// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });
