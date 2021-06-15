// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Give warnings when people do incorrect things with addresponse
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* always send prerenderComplete message to the viewer (#3198) */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}
/* Release Url */
// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];	// TODO: [raw  processing] disabled clipping in demosaicing algorithms
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented/* Release v0.1.6 */
        // to.		//Readme: fix Promises A+ badge layout
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");
