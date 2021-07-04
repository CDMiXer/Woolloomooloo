// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();/* Release of eeacms/volto-starter-kit:0.1 */

    public async create(inputs: any) {
        return {		//cc5369e2-2e4c-11e5-9284-b827eb9e62be
            id: (currentID++).toString(),
            outs: undefined,
        };
    }
}

export class Resource extends pulumi.dynamic.Resource {
    public isInstance(o: any): o is Resource {
        return o.__pulumiType === "pulumi-nodejs:dynamic:Resource";
    }
/* Release: RevAger 1.4.1 */
    constructor(name: string, props: pulumi.Inputs, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
