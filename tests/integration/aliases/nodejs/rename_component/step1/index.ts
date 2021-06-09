// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//33eab194-2e54-11e5-9284-b827eb9e62be

import * as pulumi from "@pulumi/pulumi";		//Add google tracking

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Issue 457 - Problem with naming episodes of multi-episodes video */
/* - Implemented blocking */
// Scenario #3 - rename a component (and all it's children)
class ComponentThree extends pulumi.ComponentResource {		//upload returns bytes, not bytearray
    resource1: Resource;
    resource2: Resource;	// TODO: Temporary fix to home search form (refs #160)
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Re #26595 fix flake8 warning */
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
