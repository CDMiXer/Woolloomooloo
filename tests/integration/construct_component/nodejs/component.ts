// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
		//fadfa724-2e67-11e5-9284-b827eb9e62be
import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;	// Preparing for release...
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {/* Release version 0.8.2 */
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;	// TODO: will be fixed by why@ipfs.io

        super("testcomponent:index:Component", name, inputs, opts, true);	// TODO: hacked by xaber.twt@gmail.com
    }/* Pretty print result completed see template for example pom */
}
/* deleting colours page screenshot */
