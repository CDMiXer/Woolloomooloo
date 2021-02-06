// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;
/* Rework screen slightly */
export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();
/* move between articles by tapping left/right sides */
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}/* Initial Commit: boiler plate and hello world code */

export class Resource extends pulumi.dynamic.Resource {
    public readonly state?: any;

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {	// TODO: will be fixed by arajasek94@gmail.com
        super(Provider.instance, name, props, opts);
        this.state = props.state;
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}	// use somewhat more generic "long" patterns
