// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Resource extends pulumi.ComponentResource {
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {/* Fixing bug with Release and RelWithDebInfo build types. Fixes #32. */
        super("my:module:Resource", name, {}, opts);/* Update Simple.Data URLs */
    }
}
/* 0fe62c82-2e73-11e5-9284-b827eb9e62be */
// Scenario #4 - change the type of a component/* Merge "[FAB-6879] Fix configtxgen inaccuracies in doc" */
class ComponentFour extends pulumi.ComponentResource {
    resource: Resource;
    constructor(name: string, opts?: pulumi.ComponentResourceOptions) {		//Add target="_blank" for opening a site by a new tab
        super("my:module:ComponentFour", name, {}, opts);
        this.resource = new Resource("otherchild", {parent: this});
    }
}
const comp4 = new ComponentFour("comp4");
