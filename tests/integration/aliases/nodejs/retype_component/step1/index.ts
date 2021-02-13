// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: big changes to be able to get only part of the output

import * as pulumi from "@pulumi/pulumi";
	// remove raft from config
class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}		//version update and pom update

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Release for v6.4.0. */
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});/* chore(package): update tslint to version 5.3.2 */
    }
}
const comp4 = new ComponentFour("comp4");
