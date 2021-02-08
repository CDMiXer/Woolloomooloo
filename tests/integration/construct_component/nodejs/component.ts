// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Test on Python 3.4. */

import * as pulumi from "@pulumi/pulumi";

interface ComponentArgs {/* Updated Release URL */
    echo: pulumi.Input<any>;	// weatherforecast code
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;		//Update AutoForm.php
    public readonly childId!: pulumi.Output<pulumi.ID>;/* Make gyp files comply with gyp syntax checker. */

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};/* Release 0.7.1 */
        inputs["echo"] = args.echo;		//Update mok.js
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}

