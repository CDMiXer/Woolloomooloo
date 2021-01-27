// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Update virtual-ova.md
import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {	// util api to ask if the api supports attach
        super("my:module:Resource", name, {}, opts);
    }
}

// Scenario #4 - change the type of a component/* sec-meta handling splitted */
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});	// added util class for looking up books by isbn via isbndb.com
    }
}
const comp4 = new ComponentFour("comp4");
