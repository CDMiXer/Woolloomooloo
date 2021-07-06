// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;/* Release version: 1.0.14 */

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}/* Release Notes for v00-11-pre1 */

export class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
