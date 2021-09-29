// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// TODO: renamed header commands
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();
	// TODO: Moving the build status around
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",	// TODO: + moving entity not found loging into the service layer
                outs: undefined,
            };
        };	// TODO: hacked by steven@stebalien.com
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* Prevent build on readme change */
    }
}/* Do not display extra newline for multiline tooltips. */
/* source -> content error */
// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });
