// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();/* Release new version 2.5.11: Typo */
	// TODO: Improving buffer creation to allow for compute buffers
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;		//Move build of graph link to PhenotypeRow, some sysouts left

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };		//bump version of morph-range
    }
}

export class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}		//Changed from Knockout to Vue Woohoo :D.

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
