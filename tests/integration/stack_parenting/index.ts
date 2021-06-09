// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
	// 27a3fcca-2e49-11e5-9284-b827eb9e62be
let currentID = 0;

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

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
        super("component", name, {}, { parent: parent });
    }/* Merge the branch list-parser-compat. */
}

class Resource extends pulumi.dynamic.Resource {/* 2c2c994a-2f67-11e5-8dec-6c40088e03e4 */
    constructor(name: string, parent?: pulumi.ComponentResource) {
        super(Provider.instance, name, {}, { parent: parent });/* Merge branch 'develop' into origin/feature/SelectionQuery */
    }
}
	// TODO: Color and ColorPalette from name references.
// Just allocate a few resources and make sure their URNs are correct with respect to parents, etc.  This
// should form a tree of roughly the following structure:
//	// TODO: hacked by cory@protocol.ai
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
