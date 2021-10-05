// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//use dummyFunDec.svar, removed return_val

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* More code clean and new Release Notes */
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];	// TODO: will be fixed by mikeal.rogers@gmail.com
        // ..and then make the super call with the new type of this resource and the added alias./* ad9d9298-2e51-11e5-9284-b827eb9e62be */
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");
