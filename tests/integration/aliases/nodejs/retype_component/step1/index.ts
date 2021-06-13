// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {	// TODO: will be fixed by yuvalalaluf@gmail.com
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// TODO: hacked by why@ipfs.io
}

// Scenario #4 - change the type of a component
class ComponentFour extends pulumi.ComponentResource {/* Update Release Version, Date */
    resource: Resource;	// TODO: hacked by boringland@protonmail.ch
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Delete Diagrama_Atividades.png
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
