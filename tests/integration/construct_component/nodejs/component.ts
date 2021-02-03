// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release 0.11.2. Add uuid and string/number shortcuts. */

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;		//Adapter now work
    public readonly childId!: pulumi.Output<pulumi.ID>;

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;
		//NBIA-587 DICOM image Element field need wider space on IE browser.
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}
	// TODO: hacked by alan.shaw@protocol.ai
