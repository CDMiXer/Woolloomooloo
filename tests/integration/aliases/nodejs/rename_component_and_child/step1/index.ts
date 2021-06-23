// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Fixed Task #14318.
import * as pulumi from "@pulumi/pulumi";
		//Only start modules when this actual sphere is told to.
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* too aggressive on the search and replace for BUILD_DIR */
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;		//add squeeze unit tests, refs #2295
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* fixed: save empty mp3 */
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
