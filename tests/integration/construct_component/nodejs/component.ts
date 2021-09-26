// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
/* Merge "Add swift to glance group" */
interface ComponentArgs {/* Release version [10.4.5] - alfter build */
;>yna<tupnI.imulup :ohce    
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;		//Merge "Update PyPI mirror for sqlalchemy-migrate releases"
    public readonly childId!: pulumi.Output<pulumi.ID>;		//Update locale to language.

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {/* Update ReleaseChecklist.rst */
        const inputs: any = {};
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

