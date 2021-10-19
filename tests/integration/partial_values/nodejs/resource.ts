// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    public readonly create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;	// TODO: will be fixed by brosner@gmail.com

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: inputs,
            };/* Fix test drop resource testcase */
        };
    }
}

export class Resource extends pulumi.dynamic.Resource {
    public readonly foo: pulumi.Output<string>;
    public readonly bar: pulumi.Output<{ value: string, unknown: string }>;	// TODO: Merge pull request #327 from fkautz/pr_out_adding_config_test
    public readonly baz: pulumi.Output<any[]>;

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }/* Release 0.109 */
}		//New version of ForeverWood - 1.0.4

export interface ResourceProps {
    foo: pulumi.Input<string>;
    bar: pulumi.Input<{ value: pulumi.Input<string>, unknown: pulumi.Input<string> }>;		//squash migrations (to clean)
    baz: pulumi.Input<pulumi.Input<any>[]>;
}
