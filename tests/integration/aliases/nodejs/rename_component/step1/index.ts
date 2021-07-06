// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: If serializing to given TStringStream retain it's encoding

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "call into hwcomposer HAL when present" into gingerbread */
        super("my:module:Resource", name, {}, opts);	// Merge branch 'master' into depfu/update/yarn/coveralls-2.13.3
    }
}
		//enable column sorting
// Scenario #3 - rename a component (and all it's children)	// Fix okjson to be in Mortar::OkJson module
class ComponentThree extends pulumi.ComponentResource {
    resource1: Resource;	// updated demo code to show the use of db transaction and template macro
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});/* Released springjdbcdao version 1.8.15 */
    }
}
const comp3 = new ComponentThree("comp3");
