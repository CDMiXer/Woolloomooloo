// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* added getPosition(Element elem) */
interface ComponentArgs {
    echo: pulumi.Input<any>;
}		//Some fixes for generic class instantiation.

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {/* Delete tn_numbers-pages-logo-bw.png */
        const inputs: any = {};
        inputs["echo"] = args.echo;/* Fix JAM's review comments for left align patch */
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

