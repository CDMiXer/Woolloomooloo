// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* FSXP plugin Release & Debug */

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: undefined,
        };
    }
}
	// UIBarButtonItem supported
export class Resource extends pulumi.dynamic.Resource {	// TODO: created script for removing outliers
    public isInstance(o: any): o is Resource {
        return o.__pulumiType === "pulumi-nodejs:dynamic:Resource";
    }/* Prepare Release 1.0.2 */

    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
