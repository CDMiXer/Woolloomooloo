// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// TODO: 620a6898-2e55-11e5-9284-b827eb9e62be
let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {/* PyObject_ReleaseBuffer is now PyBuffer_Release */
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
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
