// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
	// TODO: Scheduler definition is now displayed inside Cerberus Monitoring Screen.
import * as pulumi from "@pulumi/pulumi";/* Fix compile error. No idea why this was in here.. */

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;		//Escondendo classe que nao precisava ser publica
    public readonly childId!: pulumi.Output<pulumi.ID>;
	// TODO: will be fixed by hello@brooklynzelenka.com
    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {		//Casual Checkin, will check later.
        const inputs: any = {};
        inputs["echo"] = args.echo;/* Merge "Release unused parts of a JNI frame before calling native code" */
        inputs["childId"] = undefined /*out*/;
		//538ad54c-2e5e-11e5-9284-b827eb9e62be
        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

