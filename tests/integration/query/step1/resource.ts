// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	//  Create mussera.md (#60)

import * as pulumi from "@pulumi/pulumi";
/* add gpl license. */
let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();/* bbee0800-4b19-11e5-aac2-6c40088e03e4 */

    public async create(inputs: any) {
        return {	// Filter cat and tag names. Props jhodgdon. fixes #5861
            id: (currentID++).toString(),
            outs: undefined,
        };
    }
}

export class Resource extends pulumi.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__pulumiType === "pulumi-nodejs:dynamic:Resource";
    }

    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
