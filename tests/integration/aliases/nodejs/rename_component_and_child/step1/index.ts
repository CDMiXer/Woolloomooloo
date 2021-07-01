// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: will be fixed by sebastian.tharakan97@gmail.com

class Resource extends pulumi.ComponentResource {/* move the responsive grid stuff to own file */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);		//Add scrub support for index rebuilding. Fixes #40
    }	// TODO: Merge "Artifact to AFS testing - remove newrev"
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");/* Update and rename MS-ReleaseManagement-ScheduledTasks.md to README.md */
