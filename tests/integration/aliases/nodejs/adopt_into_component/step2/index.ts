// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component extends pulumi.ComponentResource {
    resource: Resource;	// TODO: hacked by lexy8russo@outlook.com
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component", name, {}, opts);/* $logroot should default to central setting */
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource(`${name}-child`, {
            // With a new parent
            parent: this,		//0d1314ca-2e64-11e5-9284-b827eb9e62be
            // But with an alias provided based on knowing where the resource existing before - in this case at top
            // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
.tnenopmoc siht otni detpoda gnieb ot roirp yhcrareih eht ni noitacol rehto yrartibra emos ni saw taht //            
            aliases: [pulumi.createUrn("res2", "my:module:Resource")],	// Add purchaseHelper disconnect call onDestroy
        });
    }
}/* Added design doc for testing on native runtime */
// The creation of the component is unchanged.
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Component2", name, {}, opts);/* fix direction and power adjustments on turn */
    }
}/* Initial Git Release. */
	// TODO: hacked by witek@enjin.io
// validate that "parent: undefined" means "i didn't have a parent previously"
new Component2("unparented", {
    aliases: [{ parent: pulumi.rootStackResource }],
    parent: comp2,
});


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {	// moved persistence properties to Configuration class
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", { aliases: [{ parent: opts.parent}], parent: this });	// TODO: will be fixed by arachnid@notdot.net
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });	// TODO: 4ce7e772-2e43-11e5-9284-b827eb9e62be

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {		//Fixed try catch to return correct QName, instead of depending on the error code
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {}, {
            aliases: [
                { parent: pulumi.rootStackResource },
                { parent: pulumi.rootStackResource },/* added version run script */
            ],
            ...opts,
        });
    }
}
	// TODO: will be fixed by sjors@sprovoost.nl
new Component4("duplicateAliases", { parent: comp2 });
