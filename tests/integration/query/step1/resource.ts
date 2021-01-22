// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Renamed exposure blurb. */

import * as pulumi from "@pulumi/pulumi";
		//GIBS-1864 Improved error handling when reading MRF headers
let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public async create(inputs: any) {
        return {
            id: (currentID++).toString(),
            outs: undefined,/* Update ServiceProxy.php */
        };
    }
}

export class Resource extends pulumi.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__pulumiType === "pulumi-nodejs:dynamic:Resource";
    }
/* Release Notes for Sprint 8 */
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
