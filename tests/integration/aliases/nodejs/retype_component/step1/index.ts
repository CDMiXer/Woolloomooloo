// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 4.1.0 */

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }
}/* Release v0.0.1beta4. */

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {		//Borrado logico de maestros
    resource: Resource;	// TODO: hacked by seth@sethvargo.com
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);/* Trying to add language selection remembering */
        this.resource = new Resource("otherchild", {parent: this});
    }/* Completed eq; added tests */
}
const comp4 = new ComponentFour("comp4");
