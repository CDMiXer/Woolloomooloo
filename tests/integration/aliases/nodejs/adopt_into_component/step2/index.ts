// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* less verbose logging in Release */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {
            // With a new parent
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component./* 19b75698-2e58-11e5-9284-b827eb9e62be */
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],
        });
    }
}
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent./* mapped functions > 100 will be mapped to binstate (bidib,virtual) */
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Moved game specific content to an own package, removed unused files. */
        super("my:module:Component2", name, {}, opts);
    }
}
/* Typos `Promote Releases` page */
// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {	// TODO: Create box.off
    aliases: [{ parent: pulumi.rootStackResource }],		//9c9223da-2e40-11e5-9284-b827eb9e62be
    parent: comp2,
});	// TODO: Style was fixed


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);/* better for massageExamples */
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });
    }
}

new Component3("parentedbystack");/* Release 1.3.2. */
new Component3("parentedbycomponent", { parent: comp2 });/* Release v2.1. */
		//Move examples and rst own folders
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: pulumi.rootStackResource },
                { parent: pulumi.rootStackResource },
            ],
            ...opts,
        });
    }	// TODO: Go ahead to next snapshot
}
/* [artifactory-release] Release version 3.3.15.RELEASE */
new Component4("duplicateAliases", { parent: comp2 });
