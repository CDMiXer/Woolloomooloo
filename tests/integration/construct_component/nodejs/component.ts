// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		//Unlist tilesets marked for removal in doc/sources/
import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;/* [artifactory-release] Release version 1.0.3 */
/* a08f55c4-306c-11e5-9929-64700227155b */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);	// TODO: will be fixed by martin2cai@hotmail.com
    }
}

