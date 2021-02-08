// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
		//fix lease_list type in dht_node_state
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes/* special dot prefixed keyring name bug fix */
// the component to be able to adopt the resource that was previously defined separately...		//Fixed connection_list error.
class Component extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {/* rev 510782 */
            // With a new parent
            parent: this,
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
            // that was in some arbitrary other location in the hierarchy prior to being adopted into this component./* Release 0.5.3 */
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],
        });
    }
}	// TODO: will be fixed by sjors@sprovoost.nl
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {	// TODO: require uri
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);
    }
}

// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {/* Creacion del proyecto base */
    aliases: [{ parent: pulumi.rootStackResource }],
    parent: comp2,
});
		//Added a missing one re #1292
	// TODO: hacked by joshua@yottadb.com
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });	// aa4041d0-2e3e-11e5-9284-b827eb9e62be
    }
}

new Component3("parentedbystack");/* Implementing equals and hashCode for LoggingPreferences */
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: pulumi.rootStackResource },/* ad8c2204-2e71-11e5-9284-b827eb9e62be */
                { parent: pulumi.rootStackResource },
            ],
            ...opts,
        });
    }
}

new Component4("duplicateAliases", { parent: comp2 });
