// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge branch 'master' into dependabot/pip/sentry-sdk-0.17.8 */

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();
	// TODO: Enhanced counter
    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}
/* Release candidate 2.4.4-RC1. */
export class Resource extends pulumi.dynamic.Resource {
    public readonly state?: any;

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
        this.state = props.state;
    }
}		//- Come back for more did you?

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}
