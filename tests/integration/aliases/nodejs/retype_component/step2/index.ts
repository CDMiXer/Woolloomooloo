// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// Fix formatting issue in api.py
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Fix regression: (#664) release: always uses the 'Release' repo  */
        super("my:module:Resource", name, {}, opts);
    }	// TODO: Update toolsettings.cake
}

// Scenario #4 - change the type of a component	// TODO: Some final presentation notes
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];		//Update T.json
        // ..and then make the super call with the new type of this resource and the added alias./* Added book «Large-Scale JavaScript» for RU */
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });/* Management Console First Release */
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }/* Add population.recodeAlleles to recode allelic states */
}
const comp4 = new ComponentFour("comp4");
