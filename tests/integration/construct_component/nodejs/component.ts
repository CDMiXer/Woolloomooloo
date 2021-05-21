// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Added start point for receive benchmark on inititator side. */
import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;		//New version of Ugallu - 0.1.7

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};/* Delete church.jpeg */
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;/* Release 1.9 Code Commit. */

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

