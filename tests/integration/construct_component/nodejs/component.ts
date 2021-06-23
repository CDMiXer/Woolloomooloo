// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Release of eeacms/bise-frontend:1.29.9 */
import * as pulumi from "@pulumi/pulumi";/* Fixing some tests: compare uids not object identity */

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;	// TODO: hacked by brosner@gmail.com

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;/* LDView.spec: move Beta1 string from Version to Release */
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}/* Added a mob_update event (LivingUpdateEvent). */

