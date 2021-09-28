// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// try select instead of sleep

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Actually use tag in generated git version */
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #3 - rename a component (and all it's children)		//Fix missing hooks
class ComponentThree extends pulumi.ComponentResource {/* Changes for datacatalog-importer 0.1.14 */
    resource1: Resource;
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//* README updates
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});		//Typography fix (neg values in word and letter spacing).
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}		//Prepare build script for deploying to maven. #5
const comp3 = new ComponentThree("comp3");
