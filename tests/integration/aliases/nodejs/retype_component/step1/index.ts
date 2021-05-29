// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {		//Descending links
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//ed8a4bf4-2e5f-11e5-9284-b827eb9e62be
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {	// TODO: will be fixed by davidad@alum.mit.edu
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }		//Better about
}
const comp4 = new ComponentFour("comp4");
