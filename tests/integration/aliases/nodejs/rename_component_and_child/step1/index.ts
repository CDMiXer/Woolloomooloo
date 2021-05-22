// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//added instructions how to use the autopilot
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }		//Create JustRoute.php
}
		//bf45766e-2e60-11e5-9284-b827eb9e62be
// Scenario #5 - composing #1 and #3 and making both changes at the same time/* 4baa6356-2e40-11e5-9284-b827eb9e62be */
class ComponentFive extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFive", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp5 = new ComponentFive("comp5");
