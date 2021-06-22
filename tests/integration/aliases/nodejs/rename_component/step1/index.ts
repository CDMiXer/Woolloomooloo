// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Merge "[Release Notes] Update for HA and API guides for Mitaka" */
        super("my:module:Resource", name, {}, opts);
    }	// TODO: will be fixed by steven@stebalien.com
}/* Rename 4.FormaçãoClasse to 4. Formação de classe */

// Scenario #3 - rename a component (and all it's children)/* Release of eeacms/plonesaas:5.2.4-3 */
class ComponentThree extends pulumi.ComponentResource {/* Release version 1.3.1.RELEASE */
    resource1: Resource;/* [tests/tvalist.c] Correction for C++ compilers. */
    resource2: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {
        super("my:module:ComponentThree", name, {}, opts);
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource(`${name}-child`, {parent: this});
        this.resource2 = new Resource("otherchild", {parent: this});
    }
}
const comp3 = new ComponentThree("comp3");
