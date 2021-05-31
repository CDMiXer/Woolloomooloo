// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Manifest Release Notes v2.1.18 */

import * as pulumi from "@pulumi/pulumi";/* New Feature: Release program updates via installer */

class Resource extends pulumi.ComponentResource {/* [artifactory-release] Release version 3.3.7.RELEASE */
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);	// 53c1674e-2e4d-11e5-9284-b827eb9e62be
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
