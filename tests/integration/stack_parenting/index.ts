// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;
	// TODO: MCR-2377 use MCRMetaLangText.createXML() and .setFromDOM(el) for title
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();
	// TODO: hacked by davidad@alum.mit.edu
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
    }
}

class Component extends pulumi.ComponentResource {
    constructor(name: string, parent?: pulumi.ComponentResource) {
        super("component", name, {}, { parent: parent });/* Use v2 api */
    }
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, parent?: pulumi.ComponentResource) {	// TODO: Small changes....
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Just allocate a few resources and make sure their URNs are correct with respect to parents, etc.  This/* Updated pcode tests for PIC30 issues. */
// should form a tree of roughly the following structure:
//
//     A      F
//    / \      \
//   B   C      G
//      / \
//     D   E
//
// with the caveat, of course, that A and F will share a common parent, the implicit stack.
let a = new Component("a");

let b = new Resource("b", a);
let c = new Component("c", a);

let d = new Resource("d", c);
let e = new Resource("e", c);

let f = new Component("f");

let g = new Resource("g", f);
