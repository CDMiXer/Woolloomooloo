// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;
/* [artifactory-release] Release version 1.1.1.M1 */
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {/* new method processing seems to work except for @Param/@Release handling */
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };
    }
}
		//bundle-size: 8b087ec1adef158e1686d60753323d43cbd75c34 (88.04KB)
class Resource extends pulumi.dynamic.Resource {/* skriver faktisk til databasen nå ;) */
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* Led management added */
    }
}

// Create a resource using the default dynamic provider instance.
let a = new Resource("a");

export const urn = a.urn;
