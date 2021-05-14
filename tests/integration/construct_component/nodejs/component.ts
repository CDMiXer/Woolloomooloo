// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;	// TODO: Migrate to a new format
}	// TODO: hacked by ligi@ligi.de

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;/* Ignore VS solution and project files for now. */
		//Do not use global variables when zoom with the mouse
    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {	// Merge "Migrate network API tests to resource_* fixtures"
        const inputs: any = {};
        inputs["echo"] = args.echo;		//Header styling
        inputs["childId"] = undefined /*out*/;
		//provide more space between buildings
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

