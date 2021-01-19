// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Updating build-info/dotnet/roslyn/dev15.7 for beta4-62729-08
        super("my:module:Resource", name, {}, opts);/* Fixed WP8 Release compile. */
    }
}/* Merge branch '3.x-dev' into feature/DTGB-626 */

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];/* Fix on delAllUserRates */
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });	// TODO: hacked by admin@multicoin.co
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented/* Simplify links in README.md */
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}	// TODO: added apigen docs
const comp4 = new ComponentFour("comp4");
