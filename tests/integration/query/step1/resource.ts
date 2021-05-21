// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async create(inputs: any) {	// TODO: will be fixed by alex.gaynor@gmail.com
        return {
            id: (currentID++).toString(),
            outs: undefined,	// Include Travis badge in README
        };
    }	// TODO: Change Getting Started Link
}

export class Resource extends pulumi.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__pulumiType === "pulumi-nodejs:dynamic:Resource";	// Use the new model in the rectangle
    }

    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);/* Initial Production version */
    }
}
