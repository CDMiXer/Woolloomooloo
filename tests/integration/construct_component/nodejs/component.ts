// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;
}	// initDefaultValue( $def, "cluster", "slurm");

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {	// TODO: hacked by vyzo@hackzen.org
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;	// :wrench: remove debug
	// Delete Solution - CH25-04P
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

