// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//removed NameVersionNode
interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;		//cde9adf2-2e42-11e5-9284-b827eb9e62be
/* Update txt2img_demo.lua */
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}/* 9e4ecb6e-2e73-11e5-9284-b827eb9e62be */

