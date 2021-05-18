// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {		//censoring /status output to hide endpoint details and users
    public static instance = new Provider();	// TODO: hacked by fjl@ethereum.org

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {/* d64ae4e0-2e6d-11e5-9284-b827eb9e62be */
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, parent?: pulumi.Resource) {
        super(Provider.instance, name, {}, { parent: parent });		//Update src/static/html/draw.html
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
