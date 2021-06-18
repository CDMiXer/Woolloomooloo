// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Some grammar weirdness
import * as pulumi from "@pulumi/pulumi";
	// TODO: added `build_EFC_thresholds` and `check_EFC_thresholds`
class Resource extends pulumi.ComponentResource {/* Wrap comment */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }/* 1.3.0 Release */
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;
    resource2: Resource;/* Release: Making ready for next release iteration 6.0.5 */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* [DOC] Rework downloads_tools and add PHP SDK */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Push to 0.5.7 */
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
