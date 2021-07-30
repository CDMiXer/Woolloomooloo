// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// updated collections to use our custom collection class
import * as pulumi from "@pulumi/pulumi";		//Update program02.c

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});/* More bug fixes for ReleaseID->ReleaseGroupID cache. */
    }
}/* Release 0.9.3 */
const comp4 = new ComponentFour("comp4");
