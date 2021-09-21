// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Widgets already have scopeStyle
        // Add an alias that references the old type of this resource...	// Modified os_utils.read_file for Python 3.5 compatibility.
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];/* MkReleases remove method implemented. Style fix. */
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });	// TODO: hacked by mikeal.rogers@gmail.com
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.		//Fix inconsistent bubble position due to displayed timestamp.
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");
