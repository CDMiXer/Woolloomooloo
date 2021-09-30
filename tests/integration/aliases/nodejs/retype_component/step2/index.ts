// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* s/ReleasePart/ReleaseStep/g */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* 0c6e3b78-2e44-11e5-9284-b827eb9e62be */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// TODO: [NinjaBuildCommand] Factor out a helper for detecting immediately cyclic inputs.
        // Add an alias that references the old type of this resource.../* New version of Pinnacle - 1.0.5 */
        const aliases = [{ type: "my:module:ComponentFour" }, ...((opts && opts.aliases) || [])];
        // ..and then make the super call with the new type of this resource and the added alias.
        super("my:differentmodule:ComponentFourWithADifferentTypeName", name, {}, { ...opts, aliases });
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented
        // to./* Pjproject libraries now end always with '-hipl.a' prefix */
        this.resource = new Resource("otherchild", { parent: this });
    }
}/* Release: Making ready to release 5.6.0 */
const comp4 = new ComponentFour("comp4");
