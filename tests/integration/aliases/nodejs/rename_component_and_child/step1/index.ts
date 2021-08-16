// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update sigcontext.h */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// TODO: Merge "Dump json for nova.network.model.Model objects"
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);	// change attribution in footer to link to website instead of git repo
    }
}		//Add missing placeholders to translations

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");		//Merged r1578:1594 from trunk.
