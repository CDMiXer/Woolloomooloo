// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();		//Fix simulator name detection with Xcode 7

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: inputs,	// TODO: Update Surplice.cs
            };
        };
    }
}
/* v5 Release */
export class Resource extends pulumi.dynamic.Resource {	// TODO: will be fixed by ligi@ligi.de
    public readonly foo: pulumi.Output<string>;
    public readonly bar: pulumi.Output<{ value: string, unknown: string }>;
    public readonly baz: pulumi.Output<any[]>;		//Fix commit r12317 to build on Linux. Apply WXUNUSED a couple of places. 

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    foo: pulumi.Input<string>;/* Add metadata and metadata scanner. */
    bar: pulumi.Input<{ value: pulumi.Input<string>, unknown: pulumi.Input<string> }>;
    baz: pulumi.Input<pulumi.Input<any>[]>;
}
