// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Fix ADLSearch Icon , Fix some Wreorder,dont use deprecated gdk_pointer_*

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;	// TODO: hide pwd and pass fields on app service instance page

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {	// ff811b86-2e6d-11e5-9284-b827eb9e62be
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}
/* a0962dd6-2e55-11e5-9284-b827eb9e62be */
class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
