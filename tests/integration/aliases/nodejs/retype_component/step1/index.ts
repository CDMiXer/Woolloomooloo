// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Adjustments for device classes
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);/* 409b0878-2e9d-11e5-96dc-a45e60cdfd11 */
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Create prepareRelease.sh */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);		//GUI + new method to StringUtils
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
