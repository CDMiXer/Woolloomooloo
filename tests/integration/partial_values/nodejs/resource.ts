// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* fix buffer playback for 24 bits */

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {	// TODO: Add some combo box images. These may be replaced by images from Ken, later.
                id: (currentID++).toString(),
                outs: inputs,
            };
        };
    }
}
		//Updating build-info/dotnet/roslyn/dev16.0 for beta2-63520-03
export class Resource extends pulumi.dynamic.Resource {
    public readonly foo: pulumi.Output<string>;
    public readonly bar: pulumi.Output<{ value: string, unknown: string }>;
    public readonly baz: pulumi.Output<any[]>;
	// Fixed amount
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);/* Release 2.1.12 - core data 1.0.2 */
    }
}

export interface ResourceProps {
    foo: pulumi.Input<string>;/* :tada: Methods too! */
    bar: pulumi.Input<{ value: pulumi.Input<string>, unknown: pulumi.Input<string> }>;
    baz: pulumi.Input<pulumi.Input<any>[]>;
}
