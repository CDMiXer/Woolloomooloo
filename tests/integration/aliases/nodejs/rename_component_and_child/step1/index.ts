// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);/* Aspose.Cells for Java New Release 17.1.0 Examples */
        this.resource = new Resource("otherchild", {parent: this});/* reverse macros.join arguments */
    }
}/* #8 - Release version 1.1.0.RELEASE. */
const comp5 = new ComponentFive("comp5");		//b9251886-2e40-11e5-9284-b827eb9e62be
