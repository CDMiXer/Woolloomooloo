// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,		//fd021ae2-2e75-11e5-9284-b827eb9e62be
            };
        };
    }
}	// TODO: will be fixed by steven@stebalien.com
/* Version 0.10.3 Release */
class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}/* Automatic changelog generation for PR #24756 [ci skip] */

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });/* Update DockerfileRelease */
