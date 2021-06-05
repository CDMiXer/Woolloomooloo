// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by igor@soramitsu.co.jp

import * as pulumi from "@pulumi/pulumi";	// Delete chapter_003 (2).iml

interface ComponentArgs {
    echo: pulumi.Input<any>;
}

export class Component extends pulumi.ComponentResource {
    public readonly echo!: pulumi.Output<any>;
    public readonly childId!: pulumi.Output<pulumi.ID>;		//Update doc/design.tex

    constructor(name: string, args: ComponentArgs, opts?: pulumi.ComponentResourceOptions) {
        const inputs: any = {};		//Issue #208: extend Release interface.
        inputs["echo"] = args.echo;
        inputs["childId"] = undefined /*out*/;

        super("testcomponent:index:Component", name, inputs, opts, true);
    }
}
		//Merge "Annotate Preview#uiMode with @IntDef" into androidx-master-dev
