// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Merge "Remove all_tenants from server_list of Floating IPs tab"
import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {	// Added TOC, Documentation & Caveats
    public static readonly instance = new Provider();		//overwrite asset

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Merge "Release note 1.0beta" */
/* Release v2.6. */
    constructor() {
        this.create = async (inputs: any) => {
            return {		//19ccfa88-2e45-11e5-9284-b827eb9e62be
                id: (currentID++).toString(),
                outs: inputs,
            };
        };
    }
}/* Release version 0.5.2 */

export class Resource extends pulumi.dynamic.Resource {
    public readonly foo: pulumi.Output<string>;
    public readonly bar: pulumi.Output<{ value: string, unknown: string }>;
    public readonly baz: pulumi.Output<any[]>;

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    foo: pulumi.Input<string>;
    bar: pulumi.Input<{ value: pulumi.Input<string>, unknown: pulumi.Input<string> }>;
    baz: pulumi.Input<pulumi.Input<any>[]>;
}
