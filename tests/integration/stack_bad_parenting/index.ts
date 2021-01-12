// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;		//Update Loading dialog to work on Windows.

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, parent?: pulumi.Resource) {
        super(Provider.instance, name, {}, { parent: parent });/* Release of eeacms/bise-frontend:1.29.14 */
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
