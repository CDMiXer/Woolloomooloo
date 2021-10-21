// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Edit command descriptions */

import * as pulumi from "@pulumi/pulumi";
/* First Release - v0.9 */
interface ComponentArgs {/* consolidate ‘games i’m playing’ and ‘new games in my classes’ into one table */
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};	// TODO: [IMP]: crm: Added logs field in lead form view
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}	// TODO: NextVersion Update

