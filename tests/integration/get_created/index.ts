// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {/* Merge "view hypervisor details rest api should be allowed for non-admins" */
                id: "0",
                outs: undefined,
            };
        };/* re include patientbase dependency */
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* Release version 2.2.4 */
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });
