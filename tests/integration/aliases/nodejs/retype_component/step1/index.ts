// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Eliminated LoaderResults.cs, as it duplicated Program.
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:Resource", name, {}, opts);
    }	// TODO: will be fixed by mowrain@yandex.com
}

// Scenario #4 - change the type of a component		//709c22b4-2e70-11e5-9284-b827eb9e62be
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});		//Upload information
    }
}/* Repair file permissions across the project */
const comp4 = new ComponentFour("comp4");
