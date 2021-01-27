// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//trailify gruppenzuordnung, fixes #3142
class Provider implements pulumi.dynamic.ResourceProvider {/* Some neo4j bug improvement */
    public static instance = new Provider();
	// TODO: Merge "[INTERNAL] Form: forward editable state to FormElement"
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {/* Changelog for #5409, #5404 & #5412 + Release date */
                id: "0",
                outs: undefined,
            };
        };
    }
}
/* Merge branch 'master' into globematerials */
class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}/* Detect props on <FormControl> `accepter` component */

// Create a resource using the default dynamic provider instance./* Release date added, version incremented. */
let a = new Resource("a");
		//Initial import of Seda 2 feature project (RCP).
// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });
