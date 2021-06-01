// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// Rename aspnetcore.md to aspnetcore.html
class Resource extends pulumi.ComponentResource {	// TODO: Add member_role and rename sequence_id
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {	// TODO: will be fixed by nagydani@epointsystem.org
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        // Add an alias that references the old type of this resource...
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to.
        this.resource = new Resource("otherchild", { parent: this });	// TODO: Add deployment documentation, and update API docs.
    }
}
const comp4 = new ComponentFour("comp4");
