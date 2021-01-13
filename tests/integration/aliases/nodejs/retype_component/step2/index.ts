// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Fixed a few optimization bugs. */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;/* Merge "Release Notes 6.0 - Minor fix for a link to bp" */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Enable Release Notes */
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];/* Do not “mark” the dom in contet script. */
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented	// TODO: hacked by julia@jvns.ca
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }	// TODO: Added example for script
}
const comp4 = new ComponentFour("comp4");
