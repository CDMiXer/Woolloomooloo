// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// 780032ea-2e55-11e5-9284-b827eb9e62be
import * as pulumi from "@pulumi/pulumi";
/* Support for NAMELIST and IMPLICIT NONE in __findLastSpecificationIndex */
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
	// TODO: Create is_leap_year.hpp
    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };/* Started post */
        };
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}/* copyright added :) */

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });
