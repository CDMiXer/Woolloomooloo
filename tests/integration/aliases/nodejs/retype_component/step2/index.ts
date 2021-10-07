// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* new appcache tpl */
	// TODO: Fix for crumb system proxy issues.
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}		//Move cover image sizes into the ResizeComponent

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;/* 6867bff4-2e3e-11e5-9284-b827eb9e62be */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...	// TODO: will be fixed by witek@enjin.io
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });
    }
}
const comp4 = new ComponentFour("comp4");	// Fixed complete session fetch
