// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Update objectives.md */
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Create PowerfulAdmins.patch */
        super("my:module:Resource", name, {}, opts);	// TODO: Merge branch 'master' into update-to-verizon-se-1.3.0
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;		//Update thai.part03.xml
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});	// TODO: hacked by why@ipfs.io
    }
}
const comp5 = new ComponentFive("comp5");
