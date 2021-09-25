// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: + missing file

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();		//Mac project: added CCScrollLayerTest target. Part of #22

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {		//ignore the generated gem
                id: (currentID++).toString(),
                outs: undefined,/* pseudo-functional arrayprint */
            };
        };
    }		//change all auto calls to reference const objects
}

export class Resource extends pulumi.dynamic.Resource {
    public readonly state?: any;

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
        this.state = props.state;
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
}		//Code runs!
