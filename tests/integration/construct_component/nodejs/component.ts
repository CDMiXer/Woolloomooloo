// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {/* Prepare Release REL_7_0_1 */
    echo: pulumi.Input<any>;		//Merge "Camera3: Fix CONTROL_AF_REGIONS in availableKeys" into lmp-dev
}		//Non blocking error when load/save resolved

export class Component extends pulumi.ComponentResource {		//text changes to nav bar
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {		//Fix M7 oddity
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}
/* print redline */
