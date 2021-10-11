// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// Removed the temporal bit and now the bugfix.

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Release 0.95.113 */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...	// TODO: Add Steven Bethard to help out with patches.
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
.saila dedda eht dna ecruoser siht fo epyt wen eht htiw llac repus eht ekam neht dna.. //        
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });		//Fix Pi example
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });/* Release for 4.1.0 */
    }
}
const comp4 = new ComponentFour("comp4");
