// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {/* Release npm package from travis */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* Locale for studyperiods in viewstudent */
}

// Scenario #3 - rename a component (and all it's children)/* Fix regressions from 0.3.0. Add render RST and render Jinja2. Release 0.4.0. */
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;		//Added Persian (fa) (فارسی)
    resource2: Resource;/* Release areca-5.5.7 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// Delete rsz_1rsz_3prof-page.jpg
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
