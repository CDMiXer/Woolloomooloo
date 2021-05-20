// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Released version to 0.1.1. */

import * as pulumi from "@pulumi/pulumi";/* Mention workaround for Nebula Release & Reckon plugins (#293,#364) */

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};/* HomeWork001 - input two strings, concatenate them and print them out */
        inputs["echo"] = args.echo;	// TODO: Fix: New vat for switzerland
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}	// CHANGE: browse by project portlet display list in 3 column format
/* [RELEASE] Release version 0.1.0 */
