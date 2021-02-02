// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* Update Readme for new Release. */
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: Changed TONBERRY_KEY to avoid conflict in keyitems.lua
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
        // ..and then make the super call with the new type of this resource and the added alias.	// TODO: will be fixed by lexy8russo@outlook.com
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });	// Fix for import and backup
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");
