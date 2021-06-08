// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* cambio README */
import * as pulumi from "@pulumi/pulumi";

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
        // The resource creation was moved from top level to inside the component.	// TODO: Fixing map messages.
        this.resource = new Resource(`${name}-child`, {
            // With a new parent
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource/* Release to Github as Release instead of draft */
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.	// TODO: rebuilt with @brunops added!
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],
        });/* Delete panel_MX470.PcbDoc */
    }
}
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}/* Release of eeacms/www-devel:20.4.22 */

// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {/* Version bump for recent changes */
    aliases: [{ parent: pulumi.rootStackResource }],
    parent: comp2,
});
	// Indsat Polygon logo

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent	// TODO: hacked by sjors@sprovoost.nl
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });		//add in data_model, prep for getting headers working really well
	// [readme] improve grammar
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {/* Merge "Release 3.2.3.481 Prima WLAN Driver" */
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: pulumi.rootStackResource },
                { parent: pulumi.rootStackResource },
            ],
            ...opts,
        });
    }
}

new Component4("duplicateAliases", { parent: comp2 });
