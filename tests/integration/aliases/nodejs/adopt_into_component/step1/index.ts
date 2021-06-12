// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Merge "Add Email.ENTERPRISE_CONTENT_LOOKUP_URI" */

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #2 - adopt a resource into a component
class Component extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Merge branch 'develop' into feature/SC-804_firstlogin_new_critical_functions
        super("my:module:Component", name, {}, opts);		//Merge branch 'release/1.1.15'
    }
}

const res2 = new Resource("res2");
const comp2 = new Component("comp2");

// Scenario 3: adopt this resource into a new parent.
class Component2 extends pulumi.ComponentResource {/* Release of eeacms/www:18.6.15 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Release of eeacms/eprtr-frontend:1.3.0-0 */
        super("my:module:Component2", name, {}, opts);
    }
}
new Component2("unparented");

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix	// TODO: hacked by greg@colvin.org
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 extends pulumi.ComponentResource {/* refactored status_text helper implementation */
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component3", name, {}, opts);
        new Component2(name + "-child", opts);
    }
}

new Component3("parentedbystack");
new Component3("parentedbycomponent", { parent: comp2 });

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 extends pulumi.ComponentResource {
    constructor(name: string, opts: pulumi.ComponentResourceOptions = {}) {
        super("my:module:Component4", name, {});
    }
}

new Component4("duplicateAliases", { parent: comp2 });
