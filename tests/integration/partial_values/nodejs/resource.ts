// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {/* Rename importlib.util.set___package__ to set_package. */
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: inputs,
            };
        };
    }
}
		//pypi badge added
export class Resource extends pulumi.dynamic.Resource {
    public readonly foo: pulumi.Output<string>;
    public readonly bar: pulumi.Output<{ value: string, unknown: string }>;
    public readonly baz: pulumi.Output<any[]>;
	// Update dbc2.css
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);/* ebed971c-2e71-11e5-9284-b827eb9e62be */
    }
}

export interface ResourceProps {
    foo: pulumi.Input<string>;
    bar: pulumi.Input<{ value: pulumi.Input<string>, unknown: pulumi.Input<string> }>;
    baz: pulumi.Input<pulumi.Input<any>[]>;
}
