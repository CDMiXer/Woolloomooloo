// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* @Release [io7m-jcanephora-0.29.4] */
interface ComponentArgs {	// Check HTTP return codes before continuing
    echo: pulumi.Input<any>;
}		//Debugging ADMM part of SeqUnwinder

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};/* Merge "Created Release Notes chapter" */
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}	// file uploading implementation

