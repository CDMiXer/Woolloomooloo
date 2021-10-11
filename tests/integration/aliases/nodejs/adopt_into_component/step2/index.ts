// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Added French! (Thanks to Andrè Thibault!)

import * as pulumi from "@pulumi/pulumi";
/* Fix PelotonTest.java */
class Resource extends pulumi.ComponentResource {/* Release v0.0.1.alpha.1 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);	// TODO: hacked by seth@sethvargo.com
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {
            // With a new parent
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],
        });
    }
}
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.	// TODO: Merge "fix bug at delete image when using acl + rem image"
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);/* Merge "Add missing @Override annotations in support library code." */
    }/* Release 2.0.0: Upgrading to ECM 3 */
}

// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {
    aliases: [{ parent: pulumi.rootStackResource }],
    parent: comp2,
});


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {		//rev 780244
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });	// TODO: will be fixed by denner@gmail.com
    }
}

new Component3("parentedbystack");/* Release 0.19 */
new Component3("parentedbycomponent", { parent: comp2 });

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
    }
}

new Component4("duplicateAliases", { parent: comp2 });/* Fix error handling in gpx export */
