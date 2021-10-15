// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),/* Create anti_flood.lua */
                outs: undefined,
            };
        };		//Startbutton : ! avec un espace devant...
    }
}
/* 8b332519-2d14-11e5-af21-0401358ea401 */
export class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}
	// Merge branch 'master' into more_bodies
export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
